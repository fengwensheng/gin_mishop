package main

import (
	"github.com/fengwensheng/gin_mishop/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	routes.IndexRouteInit(r)
	r.Run()
}
