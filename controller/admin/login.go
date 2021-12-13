package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (lc LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}
func (lc LoginController) DoLogin(c *gin.Context) {
	c.String(http.StatusOK, "admin doLogin")
}
