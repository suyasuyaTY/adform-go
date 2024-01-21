package server

import (
	"adform-go/controller"

	"github.com/gin-gonic/gin"
)


func NewRouter() (*gin.Engine, error) {
	router:= gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")
	router.GET("/", controller.Index)
	router.GET("/form", controller.NewAdForm)
	router.POST("/form", controller.CreateForm)
	return router, nil
}