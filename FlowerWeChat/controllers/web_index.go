package controllers

import (
	"github.com/astaxie/beego"
	"FlowerWeChat/controllers/wechat"
	"FlowerWeChat/models"
	"html/template"
)

type WebIndexController struct {
	beego.Controller
}

func (c *WebIndexController) Get() {
	var index models.Index
	var etiquetteFlowers [3]models.EtiquetteFlowers
	var styles [3]models.Styles
	i, _ := models.GetAllIndex(nil, nil, nil, nil, 0, 0)
	for _,v := range i{
		index = v.(models.Index)
	}
	e, _ := models.GetAllEtiquetteFlowers(nil, nil, nil, nil, 0, 0)
	for i,v := range e{
		etiquetteFlowers[i] = v.(models.EtiquetteFlowers)
	}
	s, _ := models.GetAllStyles(nil, nil, nil, nil, 0, 0)
	for i,v := range s{
		styles[i] = v.(models.Styles)
	}
	c.Data["indextitle"] = template.HTML(index.Title)
	c.Data["indeximage"] = index.ImagePath
	c.Data["userid"] = wechat.UserId
	c.Data["eflowerid1"] = etiquetteFlowers[0].Id
	c.Data["eflowertitle1"] = etiquetteFlowers[0].Title
	c.Data["eflowerimage1"] = etiquetteFlowers[0].ImagePath
	c.Data["eflowerid2"] = etiquetteFlowers[1].Id
	c.Data["eflowertitle2"] = etiquetteFlowers[1].Title
	c.Data["eflowerimage2"] = etiquetteFlowers[1].ImagePath
	c.Data["eflowerid3"] = etiquetteFlowers[2].Id
	c.Data["eflowertitle3"] = etiquetteFlowers[2].Title
	c.Data["eflowerimage3"] = etiquetteFlowers[2].ImagePath
	c.Data["styleid1"] = styles[0].Id
	c.Data["styletitle1"] = styles[0].Title
	c.Data["styleid2"] = styles[1].Id
	c.Data["styletitle2"] = styles[1].Title
	c.Data["styleid3"] = styles[2].Id
	c.Data["styletitle3"] = styles[2].Title
	c.TplName = "webview/index.html"
}