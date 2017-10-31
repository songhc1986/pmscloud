package routers

import (
	"github.com/songhc1986/pmscloud/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
