package controllers

import (
	"strconv"
	"fmt"
	"strings"
	"time"

	"blog/models"
	"blog/util"
)

//inherit baseController
type AdminController struct {
	baseController
}

//configuration information
func (c *AdminController) Config() {
	var result []*models.Config
	c.o.QueryTable(new(models.Config).TableName()).All(&result)
	options := make(map[string]string)
	mp := make(map[string]*models.Config)
	for _, v := range result {
		options[v.Name] = v.Value
		mp[v.Name] = v
	}
	if c.Ctx.Request.Method == "POST" {
		keys := []string{"url", "title", "keywords", "description", "email", "start", "qq"}
		for _, key := range keys {
			val := c.GetString(key)
			if _, ok := mp[key]; !ok {
				options[key] = val
				c.o.Insert(&models.Config{Name: key, Value: val})
			} else {
				opt := mp[key]
				if _, err := c.o.Update(&models.Config{Id: opt.Id, Name: opt.Name, Value: val}); err != nil {
					continue
				}
			}
		}
		c.History("Set data success", "")
	}
	c.Data["config"] = options
	c.TplName = c.controllerName + "/config.html"
}

//backgroud user login
func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		user := models.User{Username:username}
		c.o.Read(&user, "username")

		if user.Password == "" {
			c.History("The account dones ont exist", "")
		}

		if util.Md5(password) != strings.Trim(user.Password, "") {
			c.History("Password mistake", "")
		}

		user.LastIp = c.getClientIp()
		user.LoginCount = user.LoginCount + 1
		if _, err := c.o.Update(&user); err != nil {
			c.History("Abnormal login", "")
		} else {
			c.History("Login success", "/admin/main.html")
		}
		c.SetSession("user", user)
	}
	c.TplName = c.controllerName + "/login.html"
}
/*
c.TplName is equivalent to http.Handle(http.FlieServer())
it is used to find html
*/

func (c *AdminController) Logout() {
	c.DestroySession()
	c.History("Log out", "/admin/login.html")
}

//page
func (c *AdminController) About() {
	c.Ctx.WriteString("About")
}

//homepage
func (c *AdminController) Main() {
	c.TplName = c.controllerName + "/main.tpl"
}

//article
func (c *AdminController) Article() {
	categorys := []*models.Category{}
	c.o.QueryTable(new(models.Category).TableName()).All(&categorys)
	id, _ := c.GetInt("id")
	if id != 0 {
		post := models.Post{Id: id}
		c.o.Read(&post)
		c.Data["post"] = post
	}
	c.Data["categorys"] = categorys
	c.TplName = c.controllerName + "/_form.tpl"
}

//upload interface
func (c *AdminController) Upload() {
	f, h, err := c.GetFile("uploadname")
	result := make(map[string]interface{})
	img := ""
	if err == nil {
		exStrArr := strings.Split(h.Filename, ".")
		esStr := strings.ToLower(exStrArr[len(exStrArr) - 1])
		if exStr != "jpg" && exStr != "png" && exStr != "gif" {
			result["code"] = 1
			result["message"] = "Only upload .jpg or png format"
		}
		img = "/static/upload/" + util.UniqueId() + "." + exStr
		c.SaveToFile("upFilename". img) //save location in /static/upload
		result["code"] = 0
		result["message"] = img
	} else {
		result["code"] = 2
		result["message"] = "Abnormal upload" + err.Error()
	}
	defer f.Close()
	c.Data["json"] = result
	c.ServeJSON()
}

//Save
func (c *AdminController) Save() {
	post := models.Post{}
	post.UserId = 1
	post.Title = c.Input().Get("title")
	post.Content = c.Input().Get("content")
	post.IsTop, _ = c.GetInt8("is_top")
	post.Types, _ = c.GetInt8("types")
	post.Tags = c.Input().Get("tags")
	post.Url = c.Input().Get("url")
	post.CategoryId, _ = c.GetInt("cate_id")
	post.Info = c.Input().Get("info")
	post.Image = c.Input().Get("image")
	post.Created = time.Now()
	post.Updated = time.Now()

	id, _ := c.GetInt("id")
	if id == 0 {
		if _, err := c.o.Insert(&post); err != nil {
			c.History("insert data error" + err.Error(), "")
		} else {
			c.History("data insert success", "/admin/index.html")
		}
	} else {
		post.Id = id
		if _, err := c.o.Update(&post); err != nil {
			c.History("Update data error" + err.Error(), "")
		} else {
			c.History("Data Update success", "/admin/index/html")
		}
	}
}

func (c *AdminController) Delete() {
	id, err := strconv.Atoi(c.Input().Get("id"))
	if err != nil {
		c.History("Parameter error", "")
	} else {
		if _, err := c.o.Delete(&models.Post{Id: id}); err != nil {
			c.History("Failed to delete", "")
		} else {
			c.History("Delete success", "admin/index.html")
		}
	}
}

func (c *AdminController) Category() {
	categorys := []*models.Category{}
	c.o.QueryTable(new(models.Category).TableName()).All(&categorys)
	c.Data["categorys"] = categorys
	c.TplName = c.controllerName + "/category.tpl"
}

func (c *AdminController) Categoryadd() {
	id := c.Input().Get("id")
	if id != "" {
		intId, _ := strconv.Atoi(id)
		cate := models.Category{Id: intId}
		c.o.Read(&cate)
		c.Data["cate"] = cate
	}
	c.TplName = c.controllerName + "/category_add.tpl"
}

//handles inserting data fields
func (c *AdminController) CategorySave() {
	name := c.Input().Get("name")
	id := c.Input().Get("id")
	category := models.Category{}
	category.Name = name
	if id == "" {
		if _, err := c.o.Insert(&category); err != nil {
			c.History("insert data error", "")
		} else {
			c.History("inser data success", "admin/category.html")
		}
	} else {
		intId, err := strconv.Atoi(id)
		if err != nil {
			c.History("Parmeter error", "")
		}
		category.Id = intId
		if _, err := c.o.Update(&category); err != nil {
			c.History("Update data error", "")
		} else {
			c.History("insert data success", "/admin/category.html")
		}
	}
}

func (c *AdminController) CategoryDel() {
	id, err := strconv.Atoi(c.Input().Get("id"))
	if err != nil {
		c.History("Parmeter error", "")
    } else {
		if _, err := c.o.Delete(&models.Category{Id: id}); err != nil {
			c.History("Failed to delete", "")
		} else {
			c.History("Delete success", "/admin/category.html")
		}
	}
}