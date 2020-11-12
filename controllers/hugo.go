package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"gopkg.in/yaml.v2"
)

// Liste des répertoires et fichiers du répertoire hugo
var hugo []hugodoc

// Hugoli Liste des fichiers et répertoires
func (c *HugoController) Hugoli() {
	appid := c.Ctx.Input.Param(":app")
	dirid := c.Ctx.Input.Param(":dir")
	baseid := c.Ctx.Input.Param(":base")

	// Chargement des répertoires et fichiers
	if len(hugo) == 0 {
		hugoDirectory(c.Data["DataDir"].(string))
	}

	// Remplissage du contexte pour le template
	c.Data["DirId"] = dirid
	c.Data["BaseId"] = baseid
	c.Data["Search"] = ""
	c.Data["Records"] = hugo

	c.Ctx.Output.Cookie("from", fmt.Sprintf("/hugo/li/%s", appid))
	c.TplName = "hugo_list.html"
}

// Hugodoc table
type hugodoc struct {
	ID         int
	Path       string
	Base       string
	Dir        string
	Ext        string
	IsDir      string
	Level      int
	Title      string
	Draft      string
	Date       string
	Tags       string
	Categories string
	Content    string
}

type hugoMeta struct {
	Title      string   `yaml:"title"`
	Draft      bool     `yaml:"draft"`
	Date       string   `yaml:"date"`
	Tags       []string `yaml:"tags"`
	Categories []string `yaml:"categories"`
}

/**
 * hugoDirectory:
 * - lecture des répertoires /content et /data de foirexpo
 *
 **/
func hugoDirectory(hugoDirectory string) (err error) {

	// Lecture des répertoires et insertion d'un record par document
	var id int
	err = filepath.Walk(hugoDirectory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// On ne prend que le répertoire content
			if strings.Contains(path, hugoDirectory+"/content") {
				record := hugoFile(hugoDirectory, path, info)
				id++
				record.ID = id
				hugo = append(hugo, record)
				return nil
			}
			return nil
		})
	if err != nil {
		beego.Error(err)
	}

	return
}

func hugoFile(hugoDirectory string, pathAbsolu string, info os.FileInfo) (record hugodoc) {

	// On enlève le chemin absolu du path
	lenPrefixe := len(hugoDirectory + "/content")
	path := pathAbsolu[lenPrefixe:]
	if path == "" {
		return
	}

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

	return
}
