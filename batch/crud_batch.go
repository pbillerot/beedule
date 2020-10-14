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
	"fmt"
	"regexp"
	"time"

	"github.com/MichaelS11/go-scheduler"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pbillerot/beedule/app"
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
		err = BatchScheduler.Make(chain.Label, chain.Planif, startChain, chain)
		if err != nil {
			beego.Error("batch", chain.Label, chain.Planif, err)
			return
		}
		beego.Info("batch", "planification", chain.Label, chain.Planif)
		BatchScheduler.Start(chain.Label)
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
	chain.Etat = "RUN"
	updateChain(&chain)

	// Lecture des Jobs
	var jobs []Job
	num, err := o.QueryTable("jobs").Filter("chain_id", chain.ID).All(&jobs)
	if err != nil {
		beego.Error("batch", app.Chains.AliasDB, err)
		return
	}
	if num == 0 {
		beego.Info("batch", chain.Label, "aucun job trouvé")
		return
	}

	// Démarrage des jobs
	for _, job := range jobs {
		err = startJob(&job, &chain)
		if err != nil {
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

func startJob(job *Job, chain *Chain) (err error) {
	beego.Info("batch", chain.Label, job.Label, "Start")

	job.Etat = "RUN"
	t1 := time.Now()
	job.HeureDebut = t1.Format("2006-01-02 15:04:05")
	updateJob(job)

	switch job.Type {
	case "sql":
		err = runSQL(job, chain)
		if err != nil {
			chain.Etat = "KO"
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
