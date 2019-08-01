package routers

import (
	"crawler/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/crawler", &controllers.CrawMovieController{}, "*:CrawlMovie")
}