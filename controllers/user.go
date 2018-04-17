package controllers

import (
	"TantanDemo/models"
	"encoding/json"
	"errors"
	"strings"

	"github.com/astaxie/beego"
)

type Users struct {
	ID   string `json:"id"`
	NAME string `json:"name"`
	TYPE string `json:"type"`
}

type UserController struct {
	beego.Controller
}

func (c *UserController) GetUsers() {

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int
	var offset int

	//set fields
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	} else {
		fields = strings.Split("id,name,gender", ",")
	}

	// limit: 10 (default is 10)
	if v, err := c.GetInt("limit"); err == nil {
		limit = v
	} else {
		limit = 10
	}

	// offset: 0 default
	if v, err := c.GetInt("offset"); err == nil {
		offset = v
	} else {
		offset = 0
	}

	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	} else {
		sortby = append(sortby, "id")
	}

	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	} else {
		order = append(order, "asc")
	}

	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

func (c *UserController) PostUsers() {

	var user models.Users
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	// id, _ := c.GetInt64(":id")
	// name := c.GetString(":name")
	// gender, _ := c.GetInt8(":gender")

	// var user Users

	data, err := models.AddUser(&user)

	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = data
	}
	c.ServeJSON()
}
