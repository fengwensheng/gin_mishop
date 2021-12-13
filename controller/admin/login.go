package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (ctr LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}
func (ctr LoginController) DoLogin(c *gin.Context) {
	c.String(http.StatusOK, "admin doLogin")
}
