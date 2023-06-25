package routers

import (
	"github.com/astaxie/beego"
	v2 "github.com/imsilence/gocmdb/server/controllers/api/v2"

	"gocmdb/server/controllers"
	v1 "gocmdb/server/controllers/api/v1"
	"gocmdb/server/controllers/auth"
)

func init() {
	// 认证
	beego.AutoRouter(&auth.AuthController{})

	// 用户页面
	beego.AutoRouter(&controllers.UserPageController{})

	// 用户
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TokenController{})

	// 云平台页面
	beego.AutoRouter(&controllers.CloudPlatformPageController{})

	// 云平台
	beego.AutoRouter(&controllers.CloudPlatformController{})

	// 云主机页面
	beego.AutoRouter(&controllers.VirtualMachinePageController{})
	// 云主机
	beego.AutoRouter(&controllers.VirtualMachineController{})

	v1Namespace := beego.NewNamespace("/v1",
		beego.NSRouter("api/heartbeat/:uuid/", &v1.APIController{}, "*:Heartbeat"),
		beego.NSRouter("api/register/:uuid/", &v1.APIController{}, "*:Register"),
		beego.NSRouter("api/log/:uuid/", &v1.APIController{}, "*:Log"),
	)
	beego.AddNamespace(v1Namespace)

	v2Namespace := beego.NewNamespace("/v2",
		beego.NSRouter("api/heartbeat/:uuid/", &v2.APIController{}, "*:Heartbeat"),
		beego.NSRouter("api/register/:uuid/", &v2.APIController{}, "*:Register"),
		beego.NSRouter("api/log/:uuid/", &v2.APIController{}, "*:Log"),
	)
	beego.AddNamespace(v2Namespace)
}
