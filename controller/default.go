package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotImplemented(ctx *gin.Context) {
	msg := fmt.Sprintf("%s access to %s is not implemented yet", ctx.Request.Method, ctx.Request.URL)
	ctx.Header("Cache-Contrl", "no-cache")
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": msg})
}