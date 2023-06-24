package controllers

import (
	"gocmdb/server/controllers/auth"
	"time"
)

type TestController struct {
	auth.LoginRequiredController
}

func (c *TestController) Test() {
	c.Data["json"] = map[string]interface{}{"now": time.Now()}
	c.ServeJSON()
}
