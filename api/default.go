package api

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) showDefaultRoute(ctx *gin.Context) {
	ctx.JSON(0, "Welcome To Boin POS")
}
