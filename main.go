package main

import (
	"github.com/fengwensheng/gin_mishop/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/**/**/*")
	r.Static("/static", "./static")
	routes.AdminRouteInit(r)
	r.Run()
}
