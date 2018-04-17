package routers

import (
	"TantanDemo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users", &controllers.UserController{}, "get:GetUsers;post:PostUsers")
	beego.Router("/users/:user_id([0-9]+)/relationships", &controllers.RelationshipsController{}, "get:GetList")
	beego.Router("/users/:user_id([0-9]+)/relationships/:other_user_id([0-9]+)", &controllers.RelationshipsController{}, "put:SetRelationships")

}
