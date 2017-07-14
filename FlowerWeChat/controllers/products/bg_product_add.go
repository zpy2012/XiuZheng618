package products

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"FlowerWeChat/controllers/tools"
	"strings"
)

type BgProductAddController struct {
	beego.Controller
}

func (this *BgProductAddController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

var styles []interface{}

func (this *BgProductAddController) Get() {
	this.Data["UserName"] = username
	var sortby,order []string
	sortby = append(sortby,"Code")
	order = append(order,"desc")
	styles, _ = models.GetAllStyles(nil,nil,sortby,order,0,0)
	this.Data["styles"] = styles
	this.TplName = "bgview/productadd.html"
}

func (this *BgProductAddController) Post()  {
	var flower models.Flowers
	this.Ctx.Request.ParseMultipartForm(50)
	t1 := "0"+this.Ctx.Request.FormValue("type1")
	t2 := this.Ctx.Request.FormValue("type2")
	var type1, type2 string
	switch t1 {
	case "01": type1 = "男神/女神"
	case "02": type1 = "朋友"
	case "03": type1 = "长辈"
	case "04": type1 = "最佳情侣"
	}
	for _, v := range styles{
		style := v.(models.Styles)
		if strings.EqualFold(style.Code, t2) {
			type2 = style.Title
			break
		}
	}
	flower.Code = type1+"、"+type2
	query := make(map[string]string)
	query["Code"] = type1+"、"+type2
	flowers, err := models.GetAllFlowers(query,nil,nil,nil,0,0)
	if err != nil {
		var result models.Result
		result.Code = "201"
		result.Msg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	if len(flowers) >= 3 {
		var result models.Result
		result.Code = "201"
		result.Msg = "该编号已达到最大数量！"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	flower.Name = this.Ctx.Request.FormValue("flowername")
	flower.Description = this.Ctx.Request.FormValue("flowerdescription")
	flower.CreatedTime = tools.TimeNow().Format(beego.AppConfig.String("time"))
	if id, err := models.AddFlowers(&flower); err == nil {
		file, handler, err := this.Ctx.Request.FormFile("flowerimage")
		if err != nil {
			models.DeleteFlowers(int(id))
			var result models.Result
			result.Code = "201"
			result.Msg = err.Error()
			this.Data["json"] = result
			this.ServeJSON()
			return
		}else {
			filePath := beego.AppConfig.String("flowerpath")
			if imgpath, err := tools.SaveImage(file,handler,id,filePath);err == nil {
				flower.ImagePath = imgpath
				flower.Id = int(id)
				if err := models.UpdateFlowersById(&flower);err ==nil {
					var result models.Result
					result.Code = "200"
					result.Msg = "上传成功"
					result.JumpUrl = "/admin/products"
					this.Data["json"] = result
					this.ServeJSON()
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