package index

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"strconv"
	"FlowerWeChat/controllers/tools"
)

type BgIndexEditController struct {
	beego.Controller
}

func (this *BgIndexEditController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgIndexEditController) Get() {
	this.Data["UserName"] = username
	fmt.Printf("%v", index)
	this.Data["IndexTitle"] = index.Title
	this.Data["IndexId"] = index.Id
	this.Data["ImagePath"] = index.ImagePath
	this.TplName = "bgview/indexedit.html"
}

var index *models.Index

func (this *BgIndexEditController) Post() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	i := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(i)
	e := this.Ctx.Request.FormValue("edit")
	edit, _ := strconv.ParseBool(e)
	if edit == true {
		if in, err := models.GetIndexById(id); err == nil {
			index = in
			fmt.Fprint(this.Ctx.ResponseWriter,"成功")
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter,"失败")
		}
	}else {
		this.Ctx.Request.ParseMultipartForm(50)
		index.Title = this.Ctx.Request.FormValue("indextitle")
		file, handler, err := this.Ctx.Request.FormFile("indeximage")
		if err == nil {
			tools.DeleteImageWithPath(index.ImagePath)
			filePath := beego.AppConfig.String("indexpath")
			if imgpath, err := tools.SaveImage(file,handler,int64(index.Id),filePath);err == nil {
				index.ImagePath = imgpath
			}else {
				fmt.Fprint(this.Ctx.ResponseWriter, err)
			}
		}
		if err := models.UpdateIndexById(index);err ==nil {
			var result models.Result
			result.Code = "200"
			result.Msg = "上传成功"
			result.JumpUrl = "/admin"
			this.Data["json"] = result
			this.ServeJSON()
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter, err)
		}
	}
}