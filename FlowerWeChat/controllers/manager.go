package controllers

import (
	"FlowerWeChat/models"
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/controllers/tools"
)

// ManagerController operations for Manager
type ManagerController struct {
	beego.Controller
}

func CheckAdmins() (admins []interface{}, err error) {

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	//// fields: col1,col2,entity.col3
	//if v := c.GetString("fields"); v != "" {
	//	fields = strings.Split(v, ",")
	//}
	//// limit: 10 (default is 10)
	//if v, err := c.GetInt64("limit"); err == nil {
	//	limit = v
	//}
	//// offset: 0 (default is 0)
	//if v, err := c.GetInt64("offset"); err == nil {
	//	offset = v
	//}
	//// sortby: col1,col2
	//if v := c.GetString("sortby"); v != "" {
	//	sortby = strings.Split(v, ",")
	//}
	//// order: desc,asc
	//if v := c.GetString("order"); v != "" {
	//	order = strings.Split(v, ",")
	//}
	//// query: k:v,k:v
	//if v := c.GetString("query"); v != "" {
	//	for _, cond := range strings.Split(v, ",") {
	//		kv := strings.SplitN(cond, ":", 2)
	//		if len(kv) != 2 {
	//			c.Data["json"] = errors.New("Error: invalid query key/value pair")
	//			c.ServeJSON()
	//			return
	//		}
	//		k, v := kv[0], kv[1]
	//		query[k] = v
	//	}
	//}

	return models.GetAllManager(query, fields, sortby, order, offset, limit)
}

func Prepare ()  {
	admins, err := CheckAdmins()
	if err != nil {
		fmt.Errorf("%v",err)
	}else if len(admins) <= 0 {
		var v models.Manager
		v.Username = "admin"
		v.Password = tools.EncryptionPassWord("123456", "zpy")
		v.CreatedTime = tools.TimeNow()
		if _, err := models.AddManager(&v); err == nil {
			fmt.Println("admin添加成功！")
		}
	}
}