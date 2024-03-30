package router

import (
	"github.com/DiegoAndresMarmota/Gorga/handler"
	"github.com/gin-gonic/gin"
)

func categoryRoutes(r *gin.Engine) {

	v1 := r.Group("/ofertas/v1")
	{
		v1.GET("/ofertas-imperdibles", handler.CategoryGetHandlerOfertas1)
		v1.GET("/ofertas-exclusivas-club", handler.CategoryGetHandlerOfertas2)
		v1.GET("/productos-club", handler.CategoryGetHandlerOfertas3)
	}
}
