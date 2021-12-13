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
		//manager
		adminRoute.GET("/manager", admin.ManagerController{}.Index)
		adminRoute.GET("/manager/add", admin.ManagerController{}.Add)
		adminRoute.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRoute.GET("/manager/delete", admin.ManagerController{}.Delete)
		//banners
		adminRoute.GET("/banners", admin.BannersController{}.Index)
		adminRoute.GET("/banners/add", admin.BannersController{}.Add)
		adminRoute.GET("/banners/edit", admin.BannersController{}.Edit)
		adminRoute.GET("/banners/delete", admin.BannersController{}.Delete)
		//main
		adminRoute.GET("/", admin.MainController{}.Index)
		adminRoute.GET("/welcome", admin.MainController{}.Welcome)
	}
}
