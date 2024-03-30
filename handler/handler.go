package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CategoryGetHandlerOfertas1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET - Ofertas imperdibles",
	})
}

func CategoryGetHandlerOfertas2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET - Ofertas Exclusivas del club",
	})
}

func CategoryGetHandlerOfertas3(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET - Ofertas Productos Club",
	})
}
