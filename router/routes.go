package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func categoryRoutes(router *gin.Engine) {

	v1 := router.Group("/ofertas/v1")
	{
		v1.GET("/ofertas-imperdibles", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET - Ofertas imperdibles",
			})
		})
		v1.GET("/ofertas-exclusivas-club", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET - Ofertas Exclusivas del club",
			})
		})
		v1.GET("/productos-club", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET - Ofertas Productos Club",
			})
		})
	}
}
