package controllers

import (
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
	"gopkg.in/yaml.v2"
)

// The following request properties are provided
// to the forward-auth target endpoint as `X-Forwarded-` headers.
// | Property          | Forward-Request Header |
// |-------------------|------------------------|
// | HTTP Method       | X-Forwarded-Method     |
// | Protocol          | X-Forwarded-Proto      |
// | Host              | X-Forwarded-Host       |
// | Request URI       | X-Forwarded-Uri        |
// | Source IP-Address | X-Forwarded-For        |
// Déclarartion dans traefik middlewares.yaml
// forwardAuth:
// 		address: "http://192.168.1.76:3945/auth" # ok pour cette url
// 		# address: "http://pbillerot.freeboxos.fr/auth"
// 		trustForwardHeader: true
// 		authResponseHeaders:
// 		- "Remote-User"
// 		- "Remote-Groups"
// 		- "Remote-Name"
// 		- "Remote-Email"

// TraefikController as
type TraefikController struct {
	beego.Controller
}

// Get of TraefikController
func (c *TraefikController) Get() {
	secure(c)
}

// Head of TraefikController
func (c *TraefikController) Head() {
	secure(c)
}

// Post of TraefikController
func (c *TraefikController) Post() {
	secure(c)
}

// Structure du fichier docker/secure-traefik.yaml
type secureTraefik struct {
	WhiteURI []string
	BlackIP  []string
}

var (
	traefik = secureTraefik{}
)

func initDisabled() {
	// Chargement de secureTraefik
	filename := "./docker/secure-traefik.yaml"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		beego.Error(err)
		return
	}
	// beego.Notice(string(buf))
	err = yaml.Unmarshal(buf, &traefik)
	if err != nil {
		beego.Error(err)
		return
	}
	beego.Info(filename, "loaded")
}

func auth(c *TraefikController) {
	c.Ctx.Request.Method = c.Ctx.Request.Header.Get("X-Forwarded-Method")
	c.Ctx.Request.Host = c.Ctx.Request.Header.Get("X-Forwarded-Host")
	c.Ctx.Request.URL, _ = url.Parse(c.Ctx.Request.Header.Get("X-Forwarded-Uri"))
	ip := c.Ctx.Request.Header.Get("X-Forwarded-For")
	beego.Notice(ip, c.Ctx.Request.Method, c.Ctx.Request.Host, c.Ctx.Request.URL)
	if c.GetSession("Username") != nil {
		c.Ctx.ResponseWriter.Header().Set("Remote-User", c.GetSession("Username").(string))
		beego.Notice(c.GetSession("Username").(string))
	}
	if c.GetSession("Groups") != nil {
		c.Ctx.ResponseWriter.Header().Set("Remote-Groups", c.GetSession("Groups").(string))
		beego.Notice(c.GetSession("Groups").(string))
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
}

// curl -I -H "X-Forwarded-Uri: /bee" -H "X-Forwarded-Method: HEAD"  http://192.168.1.76:3945/traefik
func secure(c *TraefikController) {
	c.Ctx.Request.Method = c.Ctx.Request.Header.Get("X-Forwarded-Method")
	c.Ctx.Request.Host = c.Ctx.Request.Header.Get("X-Forwarded-Host")
	uri := c.Ctx.Request.Header.Get("X-Forwarded-Uri")
	c.Ctx.Request.URL, _ = url.Parse(uri)
	ip := c.Ctx.Request.Header.Get("X-Forwarded-For")
	beego.Notice(ip, c.Ctx.Request.Method, c.Ctx.Request.Host, uri)

	// URI in traefik
	for _, path := range traefik.WhiteURI {
		if strings.HasPrefix(uri, path) {
			c.Ctx.ResponseWriter.WriteHeader(200)
			return
		}
	}
	c.Ctx.ResponseWriter.WriteHeader(500)
}
