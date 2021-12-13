package routes

import (
	"github.com/fengwensheng/gin_mishop/controller"
	"github.com/gin-gonic/gin"
)

func IndexRouteInit(r *gin.Engine) {
	r.GET("/", controller.Index)
}
