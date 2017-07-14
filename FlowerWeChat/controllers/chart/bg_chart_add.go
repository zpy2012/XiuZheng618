package chart

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"FlowerWeChat/controllers/tools"
	"strings"
)

type BgChartAddController struct {
	beego.Controller
}

func (this *BgChartAddController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgChartAddController) Get() {
	this.Data["UserName"] = username
	this.TplName = "bgview/chartadd.html"
}

func (this *BgChartAddController) Post()  {
	var eFlower models.EtiquetteFlowers
	eFlowers, err := models.GetAllEtiquetteFlowers(nil,nil,nil,nil,0,0)
	if err != nil {
		var result models.Result
		result.Code = "201"
		result.Msg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	if len(eFlowers) >= 3 {
		var result models.Result
		result.Code = "201"
		result.Msg = "达到数量上限！"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}else {
		this.Ctx.Request.ParseMultipartForm(50)
		code := "0"+this.Ctx.Request.FormValue("chartcode")
		for _, v := range eFlowers{
			m := v.(models.EtiquetteFlowers)
			if strings.EqualFold(m.Code,code) {
				var result models.Result
				result.Code = "201"
				result.Msg = "该编号已存在，请选择其他编号！"
				this.Data["json"] = result
				this.ServeJSON()
				return
			}
		}
		eFlower.Title = this.Ctx.Request.FormValue("charttitle")
		eFlower.IsSelected = 1
		eFlower.Code = code
		eFlower.CreatedTime = tools.TimeNow().Format(beego.AppConfig.String("time"))
		if id, err := models.AddEtiquetteFlowers(&eFlower); err == nil {
			file, handler, err := this.Ctx.Request.FormFile("chartimage")
			if err != nil {
				models.DeleteEtiquetteFlowers(int(id))
				var result models.Result
				result.Code = "201"
				result.Msg = err.Error()
				this.Data["json"] = result
				this.ServeJSON()
				return
			}else {
				filePath := beego.AppConfig.String("chartpath")
				if imgpath, err := tools.SaveImage(file,handler,id,filePath);err == nil {
					eFlower.ImagePath = imgpath
					eFlower.Id = int(id)
					if err := models.UpdateEtiquetteFlowersById(&eFlower);err ==nil {
						var result models.Result
						result.Code = "200"
						result.Msg = "上传成功"
						result.JumpUrl = "/admin/chart"
						this.Data["json"] = result
						this.ServeJSON()
						//fmt.Fprint(this.Ctx.ResponseWriter, "/admin")
					}else {
						var result models.Result
						result.Code = "201"
						result.Msg = err.Error()
						this.Data["json"] = result
						this.ServeJSON()
						return
					}
				}else {
					var result models.Result
					result.Code = "201"
					result.Msg = err.Error()
					this.Data["json"] = result
					this.ServeJSON()
					return
				}
			}
		}else {
			var result models.Result
			result.Code = "201"
			result.Msg = err.Error()
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
	}
}