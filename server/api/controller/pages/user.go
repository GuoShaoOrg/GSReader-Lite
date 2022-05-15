package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Login(req *gin.Context) {
	req.HTML(http.StatusOK, "access/login.html", gin.H{})
}

func (ctl *Controller) Register(req *gin.Context) {
	req.HTML(http.StatusOK, "access/register.html", gin.H{})
}