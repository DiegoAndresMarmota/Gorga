package router

import (
	gin "github.com/gin-gonic/gin"
)

func InitRoute() {

	r := gin.Default()

	categoryRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
