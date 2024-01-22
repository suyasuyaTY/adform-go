package controller

import (
	"adform-go/forms"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}

func NewAdForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "NewAdForm.html", gin.H{})
}

func CreateForm(ctx *gin.Context) {
	var body forms.AdRequestForm
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errorMessages := body.Validate()
	if len(errorMessages) > 0 {
		fmt.Println(errorMessages[0])
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessages})
		return
	}
	body.Create()
	ctx.JSON(http.StatusOK, gin.H{"OK": body})
}