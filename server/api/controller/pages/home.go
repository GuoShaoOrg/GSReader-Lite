package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller)Home(req *gin.Context) {
	req.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Main website",
	})
}
