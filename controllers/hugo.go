package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"gopkg.in/yaml.v2"
)

// Liste des répertoires et fichiers du répertoire hugo
var hugo []hugodoc

// HugoLi Liste des fichiers et répertoires
func (c *HugoController) HugoLi() {
	appid := c.Ctx.Input.Param(":app")
	dirid := c.Ctx.Input.Param(":dir")
	baseid := c.Ctx.Input.Param(":base")

	// Chargement des répertoires et fichiers
	if len(hugo) == 0 {
		hugoDirectory(c, c.Data["DataDir"].(string))
	}

	// Remplissage du contexte pour le template
	c.Data["DirId"] = dirid
	c.Data["BaseId"] = baseid
	c.Data["Search"] = ""
	c.Data["Records"] = hugo

	c.Ctx.Output.Cookie("from", fmt.Sprintf("/bee/hugo/list/%s", appid))
	c.TplName = "hugo_list.html"
}

// HugoImage Liste des fichiers et répertoires
func (c *HugoController) HugoImage() {
	appid := c.Ctx.Input.Param(":app")
	keyid := c.Ctx.Input.Param(":key")

	// Recherche du record
	var record hugodoc
	for _, rec := range hugo {
		if rec.Key == keyid {
			record = rec
			break
		}
	}
	flash := beego.ReadFromRequest(&c.Controller)
	if record.Key == "" {
		beego.Error("App not found", c.GetSession("Username").(string), appid)
		flash.Error("Fichier non trouvé : %s", keyid)
		flash.Store(&c.Controller)
	}

	// Remplissage du contexte pour le template
	c.Data["Search"] = ""
	c.Data["Record"] = record

	c.Ctx.Output.Cookie("from", fmt.Sprintf("/bee/hugo/list/%s", appid))
	c.TplName = "hugo_image.html"
}

// Hugodoc table
type hugodoc struct {
	ID         int
	Key        string
	Root       string
	Path       string
	PathAbsolu string
	Base       string
	Dir        string
	Ext        string
	IsDir      int
	Level      int
	Title      string
	Draft      string
	Date       string
	Tags       string
	Categories string
	Content    string
	URL        string
	SRC        string
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
func hugoDirectory(c *HugoController, hugoDirectory string) (err error) {

	// Lecture des répertoires et insertion d'un record par document
	var id int
	err = filepath.Walk(hugoDirectory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// On ne prend que le répertoire content
			if strings.Contains(path, hugoDirectory+"/content") {
				record := hugoFile(c, hugoDirectory, path, info)
				id++
				record.ID = id
				if record.Level == 0 {
					return nil
				}
				if record.Dir[1:] == record.Base {
					record.Key = record.Base
				} else {
					record.Key = strconv.Itoa(id)
				}
				record.URL = fmt.Sprintf("%s/%d", c.Data["DataUrl"].(string), id)
				record.SRC = fmt.Sprintf("%s/content%s", c.Data["DataUrl"].(string), record.Path)
				hugo = append(hugo, record)
				// beego.Info(record.Key, record.Root, record.Path)
				return nil
			}
			return nil
		})
	if err != nil {
		beego.Error(err)
	}

	return
}

func hugoFile(c *HugoController, hugoDirectory string, pathAbsolu string, info os.FileInfo) (record hugodoc) {

	// On elève le chemin absolu du path
	lenPrefixe := len(hugoDirectory + "/content")
	path := pathAbsolu[lenPrefixe:]
	if path == "" {
		return
	}

	record.PathAbsolu = pathAbsolu
	record.Path = path // on enlève la partie hugoDirectory du chemin
	record.Dir = filepath.Dir(path)
	record.Base = filepath.Base(path)
	if info.IsDir() {
		record.IsDir = 1
		if record.Dir == "/" {
			record.Dir += record.Base
		} else {
			record.Dir += "/" + record.Base
		}
	} else {
		record.IsDir = 0
	}
	islash := strings.Index(record.Dir[1:], "/")
	if islash > 0 {
		record.Root = record.Dir[1 : islash+1]
	} else {
		record.Root = record.Dir[1:]
	}
	record.Level = strings.Count(record.Dir, "/")
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
		record.Content = string(content[:])
	}
	return
}
