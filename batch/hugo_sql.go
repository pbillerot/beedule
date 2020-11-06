package batch

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gopkg.in/yaml.v2"
)

/**
 * hugoDirectoryToSQL:
 * - lecture des répertoires /content et /data de foirexpo
 *
 **/
func hugoDirectoryToSQL(command string) (err error) {
	re := regexp.MustCompile(`hugoDirectoryToSQL\((.*),(.*),(.*)\)`)
	match := re.FindStringSubmatch(command)
	if len(match) == 0 {
		return
	}
	hugoDirectory := match[1]
	table := match[2]
	aliasDB := match[3]

	// Raz de la table
	deleteAllRecords(table, aliasDB)
	// Lecture des répertoires et insertion d'un record par document
	err = filepath.Walk(hugoDirectory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// On ne prend que le répertoire content
			if strings.Contains(path, hugoDirectory+"/content") {
				hugoFileToSQL(table, aliasDB, hugoDirectory, path, info)
				return nil
			}
			return nil
		})
	if err != nil {
		beego.Error(err)
	}

	return
}

// Hugodoc table
type Hugodoc struct {
	ID         int    `orm:"pk;auto;column(id)"`
	Path       string `orm:"column(path)"`
	Base       string `orm:"column(base)"`
	Dir        string `orm:"column(dir)"`
	Ext        string `orm:"column(ext)"`
	IsDir      string `orm:"column(isdir)"`
	Level      int    `orm:"column(level)"`
	Title      string `orm:"column(title)"`
	Draft      string `orm:"column(draft)"`
	Date       string `orm:"column(date)"`
	Tags       string `orm:"column(tags)"`
	Categories string `orm:"column(categories)"`
	Content    string `orm:"column(content)"`
}

type hugoMeta struct {
	Title      string   `yaml:"title"`
	Draft      bool     `yaml:"draft"`
	Date       string   `yaml:"date"`
	Tags       []string `yaml:"tags"`
	Categories []string `yaml:"categories"`
}

func hugoFileToSQL(table string, aliasDB string, hugoDirectory string, pathAbsolu string, info os.FileInfo) {

	// On elève le chemin absolu du path
	lenPrefixe := len(hugoDirectory + "/content")
	path := pathAbsolu[lenPrefixe:]
	if path == "" {
		return
	}

	var record Hugodoc
	record.Path = path // on enlève la partie hugoDirectory du chemin
	record.Dir = filepath.Dir(path)
	record.Base = filepath.Base(path)
	if info.IsDir() {
		record.IsDir = "1"
		if record.Dir == "/" {
			record.Dir += record.Base
		} else {
			record.Dir += "/" + record.Base
		}
	} else {
		record.IsDir = "0"
	}
	record.Ext = filepath.Ext(path)
	record.Level = strings.Count(record.Dir, "/")
	ext := filepath.Ext(path)
	if ext == ".md" {
		// lecture des metadata du fichier markdown
		content, err := ioutil.ReadFile(pathAbsolu)
		if err != nil {
			beego.Error(err)
		}
		// Extraction des meta entre les --- meta ---
		var meta hugoMeta
		err = yaml.Unmarshal(content, &meta)
		if err != nil {
			beego.Error(err)
		}
		record.Title = meta.Title
		record.Date = meta.Date
		if meta.Draft {
			record.Draft = "1"
		} else {
			record.Draft = "0"
		}
		record.Tags = strings.Join(meta.Tags, ",")
		record.Categories = strings.Join(meta.Categories, ",")
		record.Content = string(content[:])
	}
	o := orm.NewOrm()
	o.Using(aliasDB)
	_, err := o.Insert(&record)
	if err != nil {
		beego.Info(err)
	}
}

func deleteAllRecords(table string, aliasDB string) {
	o := orm.NewOrm()
	o.Using(aliasDB)
	_, err := o.Raw("delete from " + table).Exec()
	if err != nil {
		beego.Info(err)
	}
}
