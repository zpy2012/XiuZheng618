package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TJackpot618 struct {
	Id               int    `orm:"column(id);pk"`
	RedeemCode       string `orm:"column(redeemCode);size(100);null" xlsx:"抽奖码"`
	Award            string `orm:"column(award);size(50);null" xlsx:"奖项"`
	Status           int    `orm:"column(status);null"`
	Province         string `orm:"column(province);size(20);null" xlsx:"省份"`
	Area             string `orm:"column(area);size(20);null" xlsx:"城市"`
	TrugstoreName    string `orm:"column(trugstoreName);size(100);null"`
	UserName         string `orm:"column(userName);size(20);null" xlsx:"姓名"`
	Telephone        string `orm:"column(telephone);size(50);null" xlsx:"电话"`
	TrugstoreAddress string `orm:"column(trugstoreAddress);size(150);null" xlsx:"地址"`
	ZipCode          string `orm:"column(zipCode);size(20);null" xlsx:"邮编"`
	ReceiveStatus    int    `orm:"column(receiveStatus);null" xlsx:"领奖状态"`
	UpdateTime       string `orm:"column(UpdateTime);size(255);null" xlsx:"中奖时间"`
}

func (t *TJackpot618) TableName() string {
	return "t_jackpot_618"
}

func init() {
	orm.RegisterModel(new(TJackpot618))
}

// AddTJackpot618 insert a new TJackpot618 into database and returns
// last inserted Id on success.
func AddTJackpot618(m *TJackpot618) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTJackpot618ById retrieves TJackpot618 by Id. Returns error if
// Id doesn't exist
func GetTJackpot618ById(id int) (v *TJackpot618, err error) {
	o := orm.NewOrm()
	v = &TJackpot618{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetTJackpot618ById retrieves TJackpot618 by Id. Returns error if
// Id doesn't exist
func GetTJackpot618ByCode(code string) (v *TJackpot618, err error) {
	o := orm.NewOrm()
	v = &TJackpot618{RedeemCode: code}
	if err = o.Read(v, "RedeemCode"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTJackpot618 retrieves all TJackpot618 matches certain condition. Returns empty list if
// no records exist
func GetAllTJackpot618(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TJackpot618))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, v == "true" || v == "1")
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []TJackpot618
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTJackpot618 updates TJackpot618 by Id and returns error if
// the record to be updated doesn't exist
func UpdateTJackpot618ById(m *TJackpot618) (err error) {
	o := orm.NewOrm()
	v := TJackpot618{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTJackpot618 deletes TJackpot618 by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTJackpot618(id int) (err error) {
	o := orm.NewOrm()
	v := TJackpot618{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TJackpot618{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
