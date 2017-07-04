package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yinheli/qqwry"
	"strings"
)

type MainController struct {
	beego.Controller
}

type Result struct {
	Ip     string
	Contry string
	City   string
	Msg    string
}

func (c *MainController) Get() {
	q := qqwry.NewQQwry("qqwry.dat")
	ip := c.Input().Get("ip")
	UserAgent := c.Ctx.Input.Header("User-Agent")
	result := new(Result)
	if ip == "" {
		ip = c.Ctx.Input.IP()
		q.Find(ip)
		result.Msg = "您的ip地址为"
	} else {
		defer func() {
			if err := recover(); err != nil {
				beego.Error(err)
				result.Msg = "ip有错误"
				//beego.BConfig.WebConfig.AutoRender = false
				c.Data["json"] = result
				c.ServeJSON()
			}
		}()
		q.Find(ip)
		result.Msg = "您查询的ip地址为"
	}
	result.Ip = q.Ip
	result.City = q.City
	result.Contry = q.Country
	if strings.Contains(UserAgent, "curl") {
		//beego.BConfig.WebConfig.AutoRender = false
		c.Data["json"] = result
		c.ServeJSON()
	} else {
		c.TplName = "index.html"
		c.Data["Result"] = &result
	}

	//if ip == "" {
	//
	//	c.Data["Result"] = Result{
	//		Ip:     q.Ip,
	//		Contry: q.Country,
	//		City:   q.City,
	//		Msg:    "您的ip地址为",
	//	}
	//} else {
	//
	//	q.Find(ip)
	//	c.Data["Result"] = Result{
	//		Ip:     q.Ip,
	//		Contry: q.Country,
	//		City:   q.City,
	//		Msg:    "您查询的ip地址为",
	//	}
	//}

}
