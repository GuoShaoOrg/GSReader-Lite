package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctl *Controller) Login(req *gin.Context) {
	req.HTML(http.StatusOK, "login.html", gin.H{})
}
