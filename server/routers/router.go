package routers

import (
	"github.com/astaxie/beego"

	"github.com/yunixiangfeng/gocmdb/server/controllers"
	"github.com/yunixiangfeng/gocmdb/server/controllers/auth"
)

func init() {
	beego.AutoRouter(&auth.AuthController{})
	beego.AutoRouter(&controllers.TestController{})
}
