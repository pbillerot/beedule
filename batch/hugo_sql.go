package batch

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gopkg.in/yaml.v2"
)

/**
 * hugoDirectoriesToSQL:
 * - lecture des répertoires /content et /data de foirexpo
 *
 **/
func hugoDirectoriesToSQL(hugoDirectory string, table string, aliasDB string) {
	// Raz de la table
	deleteAllRecords(table, aliasDB)
	// Lecture des répertoires et insertion d'un record par document
	err := filepath.Walk(hugoDirectory,
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
}

// Hugodocument table
type Hugodocument struct {
	Path       string `orm:"pk;column(path)"`
	Base       string `orm:"column(base)"`
	Dir        string `orm:"column(dir)"`
	Ext        string `orm:"column(ext)"`
	IsDir      string `orm:"column(isdir)"`
	Title      string `orm:"column(title)"`
	Draft      string `orm:"column(draft)"`
	Date       string `orm:"column(date)"`
	Tags       string `orm:"column(tags)"`
	Categories string `orm:"column(categories)"`
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

	var record Hugodocument
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
		record.Base = ""
	} else {
		record.IsDir = "0"
	}
	record.Ext = filepath.Ext(path)
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
