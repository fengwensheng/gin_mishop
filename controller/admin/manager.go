package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (ctr ManagerController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{})
}
func (ctr ManagerController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{})
}
func (ctr ManagerController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{})
}
func (ctr ManagerController) Delete(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager/delete.html", gin.H{})
}
