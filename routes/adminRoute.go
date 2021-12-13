package routes

import (
	"github.com/fengwensheng/gin_mishop/controller/admin"
	"github.com/gin-gonic/gin"
)

func AdminRouteInit(r *gin.Engine) {
	adminRoute := r.Group("/admin")
	{
		adminRoute.GET("/login", admin.LoginController{}.Index)
		adminRoute.POST("/doLogin", admin.LoginController{}.DoLogin)
	}
}
