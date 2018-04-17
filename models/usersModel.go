package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id     int64
	Name   string
	Gender int8 // 0 for male, 1 for female
}

type RetUser struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

const (
	MALE   int8 = 0
	FEMALE int8 = 1
)

var GenderMap = map[int8]string{
	MALE:   "male",
	FEMALE: "female",
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Users))
}

func User() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Users))
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int) (retUsers []RetUser, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))

	var l []Users
	var retusrs []RetUser
	if _, err = qs.All(&l); err == nil {
		for _, v := range l {
			retusrs = append(retusrs, RetUser{v.Id, v.Name, "user"})
		}
		return retusrs, err
	}
	return nil, err
}

func AddUser(u *Users) (*Users, error) {
	o := orm.NewOrm()
	user := Users{
		Id:     u.Id,
		Name:   u.Name,
		Gender: u.Gender,
	}
	_, err := o.Insert(&user)
	if err == nil {
		return &user, err
	}
	return nil, err
}

func getUserIds() {

}

func query(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))

	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []Users
	qs = qs.OrderBy(sortFields...).RelatedSel()
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
