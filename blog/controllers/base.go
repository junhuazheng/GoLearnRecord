package controllers

import (
	"strings"
	"blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*
determine the login
access to our controller without login is not allowed
*/

type baseController struct {
	beego.Controller 
	o orm.Ormer
	controllerName string
	actionName string
}

//Prepare() verify that the user is logged in
func (p *baseController) Prepare() {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[:len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()
	
	if strings.ToLower( p.controllerName) == "admin" && strings.ToLower(p.actionName) != "login" {
		if p.GetSession("user") == nil {
			p.History("not logged in", "/admin/login")
		}
	}

	//initializes the relevant elements of the foregroud page
	if strings.ToLower( p.controllerName) == "blog" {
		p.Data["actionName"] = strings.ToLower(actionName)
		var result []*models.Config
		p.o.QueryTable(new(models.Config).TableName()).All(&result)
		configs := make(map[string]string)
		for _, v := range result {
			configs[v.Name] = v.Value
		}
		p.Data["config"] = configs
	}
}

//logical demonstration of the jump
func (p *baseController) History(msg string, url string) {
	if url == "" {
		p.Ctx.WriteString("<script>alert('"+msg+"');window.histroy.go(-1);</script>")
		p.StopRun()
	} else {
		p.Redirect(url. 302)
	}
}

//get user IP address
func (p *baseController) getClientIP() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}