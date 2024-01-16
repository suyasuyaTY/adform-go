package server

import (
	"adform-go/controller"

	"github.com/gin-gonic/gin"
)


func NewRouter() (*gin.Engine, error) {
	router:= gin.Default()
	router.GET("/", controller.NotImplemented)
	return router, nil
}