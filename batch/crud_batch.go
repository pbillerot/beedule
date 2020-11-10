package batch

/*
Utilisation du module
- https://github.com/MichaelS11/go-scheduler

	The Cron format is in the form of:

	Seconds, Minutes, Hours, Day of month, Month, Day of week, Year

	Field name     Mandatory?   Allowed values    Allowed special characters
	----------     ----------   --------------    --------------------------
	Seconds        No           0-59              * / , -
	Minutes        Yes          0-59              * / , -
	Hours          Yes          0-23              * / , -
	Day of month   Yes          1-31              * / , - L W
	Month          Yes          1-12 or JAN-DEC   * / , -
	Day of week    Yes          0-6 or SUN-SAT    * / , - L #
	Year           No           1970-2099         * / , -
*/
import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/MichaelS11/go-scheduler"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/models"
)

// BatchScheduler as
var BatchScheduler *scheduler.Scheduler

// Chain as
type Chain struct {
	ID         int    `orm:"pk;column(chain_id)"`
	Label      string `orm:"column(label)"`
	Planif     string `orm:"column(planif)"`
	Active     string `orm:"column(active)"`
	Etat       string `orm:"column(etat)"`
	HeureDebut string `orm:"column(heuredebut)"`
	HeureFin   string `orm:"column(heurefin)"`
	DureeMn    string `orm:"column(dureemn)"`
}

// Job as
type Job struct {
	ID         int    `orm:"pk;column(job_id)"`
	ChainID    int    `orm:"column(chain_id)"`
	Label      string `orm:"column(label)"`
	Sequence   int    `orm:"column(sequence)"`
	SiErreur   string `orm:"column(sierreur)"` // 1: Job suivant 2: Etape suivante 3: Arrêt chaîne
	Active     string `orm:"column(active)"`
	Type       string `orm:"column(type)"`
	Commandes  string `orm:"column(commandes)"`
	Etat       string `orm:"column(etat)"` // INI, OK, KO, RUN
	Result     string `orm:"column(result)"`
	HeureDebut string `orm:"column(heuredebut)"`
	HeureFin   string `orm:"column(heurefin)"`
	DureeMn    string `orm:"column(dureemn)"`
}

// TableName as
func (u *Chain) TableName() string {
	return "chains"
}

// TableName as
func (u *Job) TableName() string {
	return "jobs"
}

func updateChain(chain *Chain) {
	o := orm.NewOrm()
	o.Using(app.Chains.AliasDB)
	_, err := o.Update(chain, "etat", "heuredebut", "heurefin", "dureemn")
	if err != nil {
		beego.Error("batch", chain.Label, err)
	}
}
func updateJob(job *Job) {
	o := orm.NewOrm()
	o.Using(app.Jobs.AliasDB)
	_, err := o.Update(job, "result", "etat", "heuredebut", "heurefin", "dureemn")
	if err != nil {
		beego.Error("batch", job.Label, err)
	}
}

// StartBatch démarrage des chaîne batch
func StartBatch() {
	// Connexion bd
	o := orm.NewOrm()
	o.Using(app.Chains.AliasDB)

	var chains []Chain
	num, err := o.QueryTable("chains").Filter("active", "1").All(&chains)
	if err != nil {
		beego.Error("batch", app.Chains.AliasDB, err)
		return
	}
	if num == 0 {
		beego.Error("batch", "aucune chaîne trouvée")
		return
	}

	BatchScheduler = scheduler.NewScheduler()

	for _, chain := range chains {
		if chain.Planif != "" {
			err = BatchScheduler.Make(chain.Label, chain.Planif, startChain, chain)
			if err != nil {
				beego.Error("batch", chain.Label, chain.Planif, err)
				return
			}
			beego.Info("batch", "planification", chain.Label, chain.Planif)
			BatchScheduler.Start(chain.Label)
		}
	}

}

// StopBatch as
func StopBatch() {
	if BatchScheduler != nil && len(BatchScheduler.Jobs()) > 0 {
		beego.Info("batch", "stopping...", BatchScheduler.Jobs())
		BatchScheduler.StopAllWait(time.Second)
		time.Sleep(time.Duration(2) * time.Second)
		beego.Info("batch", "stopped", BatchScheduler.Jobs())
	}
}

func startChain(dataInterface interface{}) {
	chain := dataInterface.(Chain)
	beego.Info("batch", chain.Label, "Start")
	o := orm.NewOrm()
	o.Using(app.Jobs.AliasDB)

	t1 := time.Now()
	chain.HeureDebut = t1.Format("2006-01-02 15:04:05")
	chain.HeureFin = ""
	chain.DureeMn = ""
	chain.Etat = "RUN"
	updateChain(&chain)

	// Lecture des Jobs
	var jobs []Job
	num, err := o.QueryTable("jobs").Filter("chain_id", chain.ID).OrderBy("sequence").All(&jobs)
	if err != nil {
		beego.Error("batch", app.Chains.AliasDB, err)
		return
	}
	if num == 0 {
		beego.Info("batch", chain.Label, "aucun job trouvé")
		return
	}

	// Nettoyage des jobs
	for _, job := range jobs {
		job.Etat = "INI"
		job.Result = ""
		job.HeureDebut = ""
		job.HeureFin = ""
		job.DureeMn = ""
		updateJob(&job)
	}

	// Démarrage des jobs
	for _, job := range jobs {
		startJob(&job, &chain)
		if job.Etat == "KO" {
			chain.Etat = "KO"
			if job.SiErreur == "3" { // Fin chaîne
				break
			}
		}
	}

	t2 := time.Now()
	chain.HeureFin = t2.Format("2006-01-02 15:04:05")
	chain.DureeMn = fmt.Sprintf("%s", t2.Sub(t1))
	if chain.Etat == "RUN" {
		chain.Etat = "OK"
	}
	updateChain(&chain)
	beego.Info("batch", chain.Label, "End")
}

func startJob(job *Job, chain *Chain) {
	beego.Info("batch", chain.Label, job.Label, "Start")

	var err error

	job.Etat = "RUN"
	t1 := time.Now()
	job.HeureDebut = t1.Format("2006-01-02 15:04:05")
	updateJob(job)

	switch job.Type {
	case "plugin":
		out, err := RunPlugin(job.Commandes)
		if err != nil {
			job.Etat = "KO"
		}
		job.Result = out
	case "sql":
		err = runSQL(job, chain)
		if err != nil {
			job.Etat = "KO"
		}
	}
	t2 := time.Now()
	job.HeureFin = t2.Format("2006-01-02 15:04:05")
	job.DureeMn = fmt.Sprintf("%s", t2.Sub(t1))
	if job.Etat == "RUN" {
		job.Etat = "OK"
	}
	updateJob(job)
	beego.Info("batch", chain.Label, job.Label, "End")

	return
}

func runSQL(job *Job, chain *Chain) (err error) {

	// Recherche aliasDB;
	re := regexp.MustCompile(`connect (.*);`)
	match := re.FindStringSubmatch(job.Commandes)
	aliasDB := "default"
	if len(match) > 0 {
		aliasDB = match[1]
	}

	o := orm.NewOrm()
	o.Using(aliasDB)

	// Suppression ligne connect ..;
	re = regexp.MustCompile(`connect .*;`)
	sql := re.ReplaceAllString(job.Commandes, ``)
	// Exécution de la requête
	res, err := o.Raw(sql).Exec()
	if err != nil {
		beego.Error("batch", chain.Label, job.Label, err)
		job.Result = err.Error()
	} else {
		num, _ := res.RowsAffected()
		res := fmt.Sprintf("%d rows affected", num)
		beego.Info("batch", chain.Label, job.Label, res)
		job.Result = res
	}
	return
}

// TestJob Test unitaire d'un job
func TestJob(idjob int) {
	o := orm.NewOrm()
	o.Using(app.Jobs.AliasDB)
	job := Job{ID: idjob}
	err := o.Read(&job)
	if err != nil {
		beego.Error("batch", "read job", idjob, err)
		return
	}
	chain := Chain{ID: job.ChainID}
	err = o.Read(&chain)
	if err != nil {
		beego.Error("batch", "read chain", job.ChainID, err)
		return
	}
	startJob(&job, &chain)
}

// RunPlugin as fonctions appelée par les job
func RunPlugin(command string) (string, error) {
	var out string
	var err error
	beego.Info(command)
	if strings.Contains(command, "StartStopPendule()") {
		// Pas de paramètre
		// Le pendule est-il démarré ?
		if BatchScheduler == nil || len(BatchScheduler.Jobs()) == 0 {
			StartBatch()
			out = strings.Join(BatchScheduler.Jobs()[:], ",")
			err = models.CrudExec("update parameters set value = '1' where id = 'batch_etat'", "admin")
		} else {
			StopBatch()
			out = strings.Join(BatchScheduler.Jobs()[:], ",")
			err = models.CrudExec("update parameters set value = '0' where id = 'batch_etat'", "admin")
		}
	} else if strings.Contains(command, "Wait") {
		re := regexp.MustCompile(`Wait\((.*)\)`)
		match := re.FindStringSubmatch(command)
		if len(match) > 0 {
			val := match[1]
			num, er := strconv.Atoi(val)
			if er == nil {
				time.Sleep(time.Duration(num) * time.Second)
				out = "Wait ended"
			} else {
				out = er.Error()
				err = errors.New(er.Error())
			}
		}
	} else if strings.Contains(command, "StartJob") {
		re := regexp.MustCompile(`StartJob\((.*)\)`)
		match := re.FindStringSubmatch(command)
		if len(match) > 0 {
			val := match[1]
			num, er := strconv.Atoi(val)
			if er == nil {
				TestJob(num)
			} else {
				out = er.Error()
				err = errors.New(er.Error())
			}
		}
	} else if strings.Contains(command, "hugoDirectoryToSQL") {
		err = hugoDirectoryToSQL(command)
	} else if strings.Contains(command, "contentSQLToFile") {
		err = contentSQLToFile(command)
	} else if strings.Contains(command, "deleteFile") {
		err = deleteFile(command)
	} else if strings.Contains(command, "renameFile") {
		err = renameFile(command)
	} else if strings.Contains(command, "copyFile") {
		err = copyFile(command)
	} else if strings.Contains(command, "deleteDirectory") {
		err = deleteDirectory(command)
	} else if strings.Contains(command, "renameDirectory") {
		err = renameDirectory(command)
	} else if strings.Contains(command, "createDirectory") {
		err = createDirectory(command)
	} else {
		out = "Plugin inconnu"
		err = errors.New("Plugin inconnu")
	}
	return out, err
}

func contentSQLToFile(command string) (err error) {
	re := regexp.MustCompile(`contentSQLToFile\((.*),(.*),(.*),(.*),(.*),(.*)\)`)
	match := re.FindStringSubmatch(command)
	if len(match) == 0 {
		return
	}

	aliasDB := match[1]
	tableName := match[2]
	keyID := match[3]
	keyValue := match[4]
	columnName := match[5]
	pathFile := match[6]

	sql := fmt.Sprintf("select %s from %s where %s = '%s'", columnName, tableName, keyID, keyValue)
	recs, err := models.CrudSQL(sql, aliasDB)
	if err != nil {
		beego.Error(err)
	}
	var content string
	for _, rec := range recs {
		for _, val := range rec {
			if reflect.ValueOf(val).IsValid() {
				content = val.(string)
			}
		}
	}
	err = ioutil.WriteFile(pathFile, []byte(content), 0644)
	return err
}

func deleteFile(command string) (err error) {
	re := regexp.MustCompile(`deleteFile\((.*)\)`)
	match := re.FindStringSubmatch(command)
	if len(match) == 0 {
		return
	}

	path := match[1]
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return
	}

	err = os.Remove(path)
	return err
}

func deleteDirectory(command string) (err error) {
	re := regexp.MustCompile(`deleteDirectory\((.*)\)`)
	match := re.FindStringSubmatch(command)
	if len(match) == 0 {
		return
	}

	path := match[1]
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return
	}

	err = os.Remove(path)
	return err
}

func copyFile(command string) (err error) {
	re := regexp.MustCompile(`copyFile\((.*),(.*)\)`)
	match := re.FindStringSubmatch(command)
	if len(match) == 0 {
		return
	}

	pathSource := match[1]
	_, err = os.Stat(pathSource)
	if os.IsNotExist(err) {
		return
	}
	pathDest := match[2]
	_, err = os.Stat(pathDest)
	if os.IsExist(err) {
		return
	}

	data, err := ioutil.ReadFile(pathSource)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(pathDest, data, 0644)
	return
}

func renameFile(command string) (err error) {
	re := regexp.MustCompile(`renameFile\((.*),(.*)\)`)
	match := re.FindStringSubmatch(command)
	if len(match) == 0 {
		return
	}

	pathSource := match[1]
	_, err = os.Stat(pathSource)
	if os.IsNotExist(err) {
		return
	}
	pathDest := match[2]
	_, err = os.Stat(pathDest)
	if os.IsExist(err) {
		return
	}
	err = os.Rename(pathSource, pathDest)
	return
}

func renameDirectory(command string) (err error) {
	re := regexp.MustCompile(`renameDirectory\((.*),(.*)\)`)
	match := re.FindStringSubmatch(command)
	if len(match) == 0 {
		return
	}

	pathSource := match[1]
	_, err = os.Stat(pathSource)
	if os.IsNotExist(err) {
		return
	}
	pathDest := match[2]
	_, err = os.Stat(pathDest)
	if os.IsExist(err) {
		return
	}
	err = os.Rename(pathSource, pathDest)
	return
}

func createDirectory(command string) (err error) {
	re := regexp.MustCompile(`createDirectory\((.*)\)`)
	match := re.FindStringSubmatch(command)
	if len(match) == 0 {
		return
	}

	path := match[1]
	_, err = os.Stat(path)
	if os.IsExist(err) {
		return
	}

	err = os.MkdirAll(path, 0666)
	return err
}

func runHugoServer(command string) (err error) {
	cmd := exec.Command("hugo", "server")
	cmd.Dir = command
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	return
}

func saveImage(command string) (err error) {
	re := regexp.MustCompile(`saveImage\((.*),(.*)\)`)
	match := re.FindStringSubmatch(command)
	if len(match) == 0 {
		return
	}

	path := match[1]
	sbuf := match[2]
	err = ioutil.WriteFile(path, []byte(sbuf), 0644)
	return
}
