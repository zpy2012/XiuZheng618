package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Manager struct {
	Id          int       `orm:"column(id);auto"`
	Username    string    `orm:"column(username);size(255)"`
	Password    string    `orm:"column(password);size(255)"`
	CreatedTime time.Time `orm:"column(created_time);type(datetime)"`
	UpdateTime  time.Time `orm:"column(update_time);type(timestamp)"`
}

func (t *Manager) TableName() string {
	return "manager"
}

func init() {
	orm.RegisterModel(new(Manager))
}

// AddManager insert a new Manager into database and returns
// last inserted Id on success.
func AddManager(m *Manager) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetManagerById retrieves Manager by Id. Returns error if
// Id doesn't exist
func GetManagerById(id int) (v *Manager, err error) {
	o := orm.NewOrm()
	v = &Manager{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetManagerById retrieves Manager by Id. Returns error if
// Id doesn't exist
func GetManagerByName(username string) (v *Manager, err error) {
	o := orm.NewOrm()
	v = &Manager{Username: username}
	if err = o.Read(v, "username"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllManager retrieves all Manager matches certain condition. Returns empty list if
// no records exist
func GetAllManager(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Manager))
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

	var l []Manager
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

// UpdateManager updates Manager by Id and returns error if
// the record to be updated doesn't exist
func UpdateManagerById(m *Manager) (err error) {
	o := orm.NewOrm()
	v := Manager{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteManager deletes Manager by Id and returns error if
// the record to be deleted doesn't exist
func DeleteManager(id int) (err error) {
	o := orm.NewOrm()
	v := Manager{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Manager{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
