package bg_controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"XiuZheng618/models"
	"math"
	"strconv"
	"strings"
)

type BgDetailController struct {
	beego.Controller
}

const prepage = 50

func (this *BgDetailController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (c *BgDetailController) Get() {
	c.Data["UserName"] = username

	c.ParseForm(c.Ctx.Input.RequestBody)
	p := c.Ctx.Request.FormValue("page")
	page, _ := strconv.Atoi(p)
	city := c.Ctx.Request.FormValue("city")
	offset := (page-1)*prepage

	var order, sortBy  []string
	order = append(order, "desc")
	sortBy = append(sortBy, "UpdateTime")

	query := make(map[string]string)
	query["Status"] = "1"

	if strings.EqualFold(city,"全部") {
		delete(query, "Area")
	}else {
		query["Province"] = city
	}

	r, _ := models.GetAllTJackpot618(query, nil, sortBy, order, 0, 0)

	if record, err := models.GetAllTJackpot618(query, nil, sortBy, order, int64(offset), prepage); err == nil {
		res := Paginator(page, prepage, int64(len(r)))
		c.Data["paginator"] = res
		c.Data["jilu"] = record
	}else {
		res := Paginator(page, prepage, 0)
		c.Data["paginator"] = res
	}
	c.Data["city"] = city
	c.TplName = "bgview/detail.html"
}

func (this *BgDetailController) Post() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	i := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(i)
	result, _ := models.GetTJackpot618ById(id)
	result.ReceiveStatus = 1
	if err := models.UpdateTJackpot618ById(result); err == nil {
		fmt.Fprint(this.Ctx.ResponseWriter, "成功")
	}else {
		fmt.Fprint(this.Ctx.ResponseWriter, "失败")
	}
}

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page, prepage int, nums int64) map[string]interface{} {

	//var firstpage int //前一页地址
	//var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		//firstpage = page - 1
		//lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		//firstpage = page - 3
		for i := range pages {
			pages[i] = start + i
		}
		//firstpage = page - 1
		//lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i := range pages {
			pages[i] = i + 1
		}
		//firstpage = int(math.Max(float64(1), float64(page-1)))
		//lastpage = page + 1
		//fmt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = 1
	paginatorMap["lastpage"] = totalpages
	paginatorMap["currpage"] = page
	return paginatorMap
}