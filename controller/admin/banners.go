package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BannersController struct {
}

func (ctr BannersController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/banners/index.html", gin.H{})
}
func (ctr BannersController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/banners/add.html", gin.H{})
}
func (ctr BannersController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/banners/edit.html", gin.H{})
}
func (ctr BannersController) Delete(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/banners/delete.html", gin.H{})
}
