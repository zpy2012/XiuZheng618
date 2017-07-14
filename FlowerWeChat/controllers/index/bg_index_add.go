package index

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"FlowerWeChat/controllers/tools"
)

type BgIndexAddController struct {
	beego.Controller
}

func (this *BgIndexAddController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgIndexAddController) Get() {
	this.Data["UserName"] = username
	this.TplName = "bgview/indexadd.html"
}

func (this *BgIndexAddController) Post()  {
	var index models.Index
	indexs, err := models.GetAllIndex(nil,nil,nil,nil,0,0)
	if err != nil {
		this.Ctx.ResponseWriter.Status = 0
		fmt.Fprint(this.Ctx.ResponseWriter, err)
		return
	}
	if len(indexs) >= 1 {
		this.Ctx.ResponseWriter.Status = 0
		fmt.Fprint(this.Ctx.ResponseWriter, "达到数量上限！")
		return
	}else {
		this.Ctx.Request.ParseMultipartForm(50)
		index.Title = this.Ctx.Request.FormValue("indextitle")
		index.IsSelected = 1
		index.CreatedTime = tools.TimeNow().Format(beego.AppConfig.String("time"))
		if id, err := models.AddIndex(&index); err == nil {
			file, handler, err := this.Ctx.Request.FormFile("indeximage")
			if err != nil {
				fmt.Fprint(this.Ctx.ResponseWriter, err)
			}else {
				filePath := beego.AppConfig.String("indexpath")
				if imgpath, err := tools.SaveImage(file,handler,id, filePath);err == nil {
					index.ImagePath = imgpath
					index.Id = int(id)
					if err := models.UpdateIndexById(&index);err ==nil {
						var result models.Result
						result.Code = "200"
						result.Msg = "上传成功"
						result.JumpUrl = "/admin"
						this.Data["json"] = result
						this.ServeJSON()
						//fmt.Fprint(this.Ctx.ResponseWriter, "/admin")
					}else {
						fmt.Fprint(this.Ctx.ResponseWriter, err)
					}
				}else {
					fmt.Fprint(this.Ctx.ResponseWriter, err)
				}
			}
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter, err)
		}
	}
}