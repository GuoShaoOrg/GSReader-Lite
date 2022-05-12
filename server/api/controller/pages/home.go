package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Home(req *gin.Context) {
	req.HTML(http.StatusOK, "index.html", gin.H{
		"username": "管理员",
	})
}

func (ctl *Controller) HomeContainerTmpl(req *gin.Context) {
	req.HTML(http.StatusOK, "home-container.html", gin.H{
	})
}
