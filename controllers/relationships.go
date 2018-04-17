package controllers

import (
	"TantanDemo/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type RelationshipsController struct {
	beego.Controller
}

// var relation = [2]Relationships{{"3123", "liked", "relationship"}, {"2133", "disliked", "relationships"}}

func (c *RelationshipsController) GetList() {
	id, _ := c.GetInt64(":user_id", 0)
	data, err := models.GetRelationshipsById(id)

	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = data
	}
	c.ServeJSON()
}

func (c *RelationshipsController) SetRelationships() {
	userId, _ := c.GetInt64(":user_id", 0)

	// state := c.GetString(":state", "")
	// stateMap := models.StateRevMaps[state]

	var relation models.RetRelationships
	json.Unmarshal(c.Ctx.Input.RequestBody, &relation)
	data, err := models.SetRelationships(userId, relation)

	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = data
	}
	c.ServeJSON()
}
