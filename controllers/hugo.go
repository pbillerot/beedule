package controllers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/yaml.v2"
)

// Liste des répertoires et fichiers du répertoire hugo
var hugo []hugodoc
var hugoRacine string

// HugoList Liste des fichiers et répertoires
func (c *HugoController) HugoList() {
	appid := c.Ctx.Input.Param(":app")
	dirid := c.Ctx.Input.Param(":dir")
	baseid := c.Ctx.Input.Param(":base")

	// Chargement des répertoires et fichiers
	if len(hugo) == 0 {
		hugoRacine = c.Data["DataDir"].(string) + "/content"
		hugoDirectoryRecord(c, c.Data["DataDir"].(string))
	}

	// Remplissage du contexte pour le template
	c.Data["DirId"] = dirid
	c.Data["BaseId"] = baseid
	c.Data["Search"] = ""
	c.Data["Records"] = hugo

	c.Ctx.Output.Cookie("from", fmt.Sprintf("/bee/hugo/list/%s", appid))
	c.TplName = "hugo_list.html"
}

// HugoImage Visualiser Modifier une image
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

	if c.Ctx.Input.Method() == "POST" {
		// ENREGISTREMENT DE L'IMAGE
		simage := c.GetString("image")
		b64data := simage[strings.IndexByte(simage, ',')+1:]
		unbased, err := base64.StdEncoding.DecodeString(b64data)
		// img, _, err := image.Decode(bytes.NewReader([]byte(element.SQLout)))
		if err != nil {
			msg := fmt.Sprintf("HugoImage %s : %s", record.PathAbsolu, err)
			beego.Error(msg)
			flash.Error(msg)
			flash.Store(&c.Controller)
			ReturnFrom(c.Controller)
		}

		outputFile, err := os.Create(record.PathAbsolu)
		if err != nil {
			msg := fmt.Sprintf("HugoImage %s : %s", record.PathAbsolu, err)
			beego.Error(msg)
			flash.Error(msg)
			flash.Store(&c.Controller)
			ReturnFrom(c.Controller)
		}
		defer outputFile.Close()

		outputFile.Write(unbased)
		// Fermeture de la fenêtre
		c.TplName = "bee_close.html"
		return
	}

	// Remplissage du contexte pour le template
	c.Data["Record"] = record
	c.Data["KeyID"] = keyid
	c.Ctx.Output.Cookie("hugo-"+appid, keyid)
	c.Ctx.Output.Cookie("from", fmt.Sprintf("/bee/hugo/image/%s", appid))
	c.TplName = "hugo_image.html"
}

// HugoDocument Visualiser Modifier un document
func (c *HugoController) HugoDocument() {
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

	if c.Ctx.Input.Method() == "POST" {
		// ENREGISTREMENT DU DOCUMENT
		document := c.GetString("document")
		if record.PathReal != "" {
			err = ioutil.WriteFile(record.PathReal, []byte(document), 0644)
			err = ioutil.WriteFile(record.PathAbsolu, []byte(document), 0644)
		} else {
			err = ioutil.WriteFile(record.PathAbsolu, []byte(document), 0644)
		}
		if err != nil {
			msg := fmt.Sprintf("HugoImage %s : %s", record.PathAbsolu, err)
			beego.Error(msg)
			flash.Error(msg)
			flash.Store(&c.Controller)
			ReturnFrom(c.Controller)
		}

		// Vidage de Hugo puis reconstruction
		hugo = nil
		hugoDirectoryRecord(c, c.Data["DataDir"].(string))
		for _, rec := range hugo {
			if rec.Key == keyid {
				record = rec
				break
			}
		}
		// Demande d'actualisation de l'arborescence
		c.Ctx.Output.Cookie("hugo-refresh-"+appid, "true")

		// Fermeture de la fenêtre
		// c.TplName = "bee_parent.html"
		// return

	}

	// Remplissage du contexte pour le template
	c.Data["Record"] = record
	c.Data["KeyID"] = keyid
	c.Ctx.Output.Cookie("hugo-"+appid, keyid)
	c.Ctx.Output.Cookie("from", fmt.Sprintf("/bee/hugo/document/%s", appid))
	c.TplName = "hugo_document.html"
}

// HugoFile Gestion du fichier
func (c *HugoController) HugoFile() {
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
	c.Data["Record"] = record
	c.Data["KeyID"] = keyid

	c.Ctx.Output.Cookie("from", fmt.Sprintf("/bee/hugo/file/%s/%s", appid, keyid))
	c.TplName = "hugo_file.html"
}

// HugoDirectory Gestion du répertoire
func (c *HugoController) HugoDirectory() {
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
	c.Data["Record"] = record
	c.Data["KeyID"] = keyid

	c.Ctx.Output.Cookie("from", fmt.Sprintf("/bee/hugo/directory/%s/%s", appid, keyid))
	c.TplName = "hugo_directory.html"
}

// HugoFileMv Renommer le fichier
func (c *HugoController) HugoFileMv() {
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
		ReturnFrom(c.Controller)
	}

	newFile := c.GetString("new_name")
	err = os.Rename(record.PathAbsolu, hugoRacine+"/"+newFile)
	if err != nil {
		msg := fmt.Sprintf("HugoImage %s : %s", newFile, err)
		beego.Error(msg)
		flash.Error(msg)
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
	}

	// Le cookie ancrage est déplacé sur le répertoire root
	c.Ctx.Output.Cookie("hugo-"+appid, record.Root)

	// Vidage de Hugo pour reconstruction
	hugo = nil

	// Fermeture de la fenêtre
	c.TplName = "bee_parent.html"
	return

}

// HugoFileCp Recopier le fichier ou répertoire
func (c *HugoController) HugoFileCp() {
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

	newFile := c.GetString("copy_file")
	data, err := ioutil.ReadFile(record.PathAbsolu)
	if err != nil {
		msg := fmt.Sprintf("HugoImage %s : %s", record.Path, err)
		beego.Error(msg)
		flash.Error(msg)
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
	}
	err = ioutil.WriteFile(hugoRacine+"/"+newFile, data, 0755)
	if err != nil {
		msg := fmt.Sprintf("HugoImage %s : %s", newFile, err)
		beego.Error(msg)
		flash.Error(msg)
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
	}

	// Vidage de Hugo pour reconstruction
	hugo = nil

	// Fermeture de la fenêtre
	c.TplName = "bee_parent.html"
	return
}

// HugoFileRm Supprimer le fichier ou répertoire
func (c *HugoController) HugoFileRm() {
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

	err = os.RemoveAll(record.PathAbsolu)
	if err != nil {
		msg := fmt.Sprintf("HugoImage %s : %s", record.Path, err)
		beego.Error(msg)
		flash.Error(msg)
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
	}

	// Le cookie ancrage est déplacé sur le répertoire root
	c.Ctx.Output.Cookie("hugo-"+appid, record.Root)

	// Vidage de Hugo pour reconstruction
	hugo = nil

	// Fermeture de la fenêtre
	c.TplName = "bee_parent.html"
	return
}

// HugoFileUpload Charger le fichier sur le serveur
func (c *HugoController) HugoFileUpload() {
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

	file, handler, err := c.Ctx.Request.FormFile("new_file")
	if err != nil {
		msg := fmt.Sprintf("HugoDirectory %s : %s", "new_file", err)
		beego.Error(msg)
		flash.Error(msg)
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		msg := fmt.Sprintf("HugoDirectory %s : %s", handler.Filename, err)
		beego.Error(msg)
		flash.Error(msg)
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
	}
	filepath := fmt.Sprintf("%s/%s", record.PathAbsolu, handler.Filename)
	err = ioutil.WriteFile(filepath, fileBytes, 0755)
	if err != nil {
		msg := fmt.Sprintf("HugoDirectory %s : %s", handler.Filename, err)
		beego.Error(msg)
		flash.Error(msg)
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
	}

	// Le cookie ancrage est déplacé sur le répertoire root
	c.Ctx.Output.Cookie("hugo-"+appid, record.Root)

	// Vidage de Hugo pour reconstruction
	hugo = nil

	// Fermeture de la fenêtre
	c.TplName = "bee_parent.html"
	return
}

// HugoFileMkdir Créer un répertoire
func (c *HugoController) HugoFileMkdir() {
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

	newDir := c.GetString("new_dir")
	err = os.MkdirAll(hugoRacine+"/"+newDir, 0755)
	if err != nil {
		msg := fmt.Sprintf("HugoImage %s : %s", newDir, err)
		beego.Error(msg)
		flash.Error(msg)
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
	}

	// Vidage de Hugo pour reconstruction
	hugo = nil

	// Fermeture de la fenêtre
	c.TplName = "bee_parent.html"
	return
}

// Hugodoc table
type hugodoc struct {
	ID          int
	Key         string
	Root        string
	Path        string
	Prefix      string
	PathAbsolu  string
	PathReal    string // path réel du fichier Ex. /data/config.yaml
	Base        string
	Dir         string
	Ext         string
	IsDir       int
	Level       int
	Title       string
	Draft       string
	Date        string
	DatePublish string
	DateExpiry  string
	Inline      bool
	Tags        string
	Categories  string
	Content     string
	HugoPath    string // url de la page sur le site
	URL         string
	SRC         string
}

type hugoMeta struct {
	Title       string   `yaml:"title"`
	Draft       bool     `yaml:"draft"`
	Date        string   `yaml:"date"`
	DatePublish string   `yaml:"publishDate"`
	DateExpiry  string   `yaml:"expiryDate"`
	Tags        []string `yaml:"tags"`
	Categories  []string `yaml:"categories"`
}

var hugoLinks = []map[string]string{
	{
		"title":  "Site de test",
		"url":    "http://localhost:1313",
		"posx":   "left",
		"target": "hugo_test",
	},
}

/**
 * hugoDirectoryRecord:
 * - lecture des répertoires /content et /data de foirexpo
 *
 **/
func hugoDirectoryRecord(c *HugoController, hugoDirectory string) (err error) {

	// Lecture des répertoires et insertion d'un record par document
	var id int
	err = filepath.Walk(hugoDirectory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// On ne prend que le répertoire content
			if strings.Contains(path, hugoRacine) {
				record := hugoFileRecord(c, hugoDirectory, path, info)
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

func hugoFileRecord(c *HugoController, hugoDirectory string, pathAbsolu string, info os.FileInfo) (record hugodoc) {

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
	if ext == ".yaml" {
		// lecture du fichier yaml
		content, err := ioutil.ReadFile(pathAbsolu)
		if err != nil {
			beego.Error(err)
		}
		record.Content = string(content[:])
		// le fichier a son clone dans /data
		record.PathReal = strings.Replace(record.PathAbsolu, "/content/site/", "/data/", 1)
	}
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
		record.DatePublish = meta.DatePublish
		record.DateExpiry = meta.DateExpiry
		record.Inline = true
		if record.DatePublish != "" && record.DatePublish > time.Now().Format("2006-01-02") {
			record.Inline = false
		}
		if record.DateExpiry != "" && record.DateExpiry <= time.Now().Format("2006-01-02") {
			record.Inline = false
		}
		if meta.Draft {
			record.Draft = "1"
			record.Inline = false
		} else {
			record.Draft = "0"
		}
		record.Tags = strings.Join(meta.Tags, ",")
		record.Categories = strings.Join(meta.Categories, ",")
		record.Content = string(content[:])
		record.HugoPath = strings.Replace(record.Path, ".md", "", 1)
	}
	return
}
