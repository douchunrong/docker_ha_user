package routers

import (
	"USC-HA-User/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hauser", &controllers.HauserController{})
	beego.Router("/hauser/:id", &controllers.HauserController{})
}
