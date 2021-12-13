package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (ctr MainController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
}

func (ctr MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
