package controllers

import (
	"time"
	"crawler/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type CrawMovieController struct {
	beego.Controller
}

func (c *CrawMovieController) CrawMovie {
	//利用结构体变量定义一个变量，然后调用方法，用内部变量存储获取的信息
	var movieInfo models.MovieInfo
	//连接redis
	models.ConnectRedis("127.0.0.1:6379")

	//爬虫入口
	sUrl := "http://movie.douban.com/subject/26985127"

	//将它加入redis的队列
	models.PutinQueue(sUrl)

	for {
		//获取队列长度，第二次执行此循环时不会为空，下方有一个获取的代码段
		length := models.GetQueueLength()
		if length == 0 {
			break //如果url队列为空，则退出当前循环
		}

		//从队列获取url
		sUrl = models.ProfromQueue()

		//判断电影的url是否被访问
		if models.IsVisit(sUrl) {
			continue
		}

		req := httplib.Get(sUrl)

		//获取页面信息
		sMovieHtml, err := req.String()
		if err != nil {
			panic(err)
		}

		//通过获取名字来判断它是不是电影
		movieInfo.Movie_name = models.GetMovieName(sMovieHtml)

		//如果它是电影，则获取其他信息，并加入mysql数据库
		if movieInfo.Movie_name != "" {
			movieInfo.Movie_director = models.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_main_character = models.GetMovieMainCharacters(sMovieHtml)
			movieInfo.Movie_type = models.GetMovieGenre(sMovieHtml)
			movieInfo.Movie_grade = models.GetMovieGrade(sMovieHtml)
			movieInfo.Movie_on_time = models.GetMovieOnTime(sMovieHtml)
			movieInfo.Movie_span = models.GetMovieRunningTime(sMovieHtml)
			movieInfo.Movie_writer = models.GetMovieWriter(sMovieHtml)
			movieInfo.Movie_country = models.GetMovieCountry(sMovieHtml)

			//加入数据库
			models.AddMovie(&movieInfo)
		}
		//提取该页面的所有连接
		//urls为字符串数组类型
		urls := models.GetMovieUrls(sMovieHtml) //urls变量用于存取调用获取的电影地址
		for _, url := range urls {

			//url进入队列
			models.PutinQueue(url)
			//爬去结束打印提示
			c.Ctx.WriteString("<br>" + url + "</br>")

			//遍历字符串数组将每一个电影地址加入队列，执行程序之前，先启动redis(redis-server)
			//之后连接到redis(redis-cli),在浏览器输入网址，在命令行输入keys * 发现多了一个url_queue 执行lrange url_queue 0 -1

		}

		//sUrl记录到访问set中，表明已经放问过了
		models.AddToSet(sUrl)

		//为了防止爬取速度过快，每次爬完一部电影休息一秒
		time.Sleep(time.Second)
	}
	
	c.Ctx.WriteString("end of crawler!")
}