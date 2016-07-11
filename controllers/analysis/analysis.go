package analysis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowData(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "analysis.tmpl", gin.H{
		"title": "Analysis Page",
	})
}
