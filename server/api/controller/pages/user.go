package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Login(req *gin.Context) {
	req.HTML(http.StatusOK, "page/login.html", gin.H{})
}

func (ctl *Controller) Register(req *gin.Context) {
	req.HTML(http.StatusOK, "page/register.html", gin.H{})
}