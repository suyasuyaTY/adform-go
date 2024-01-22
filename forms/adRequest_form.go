package forms

import (
	"adform-go/models"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type AdRequestForm struct {
	Name        string `json:"name" validate:"required"`
    Mail		string `json:"mail" validate:"required,mail"`
	MessageJa	string `json:"messageJa" validate:"required"`
	MessageEn	string `json:"messageEn" validate:"required"`
	URLJa		string `json:"urlJa" validate:"required,url"`
	URLEn		string `json:"urlEn" validate:"required,url"`
	StartDate	string `json:"start-date" validate:"required,date"`
	EndDate		string `json:"end-date" validate:"required,date"`

}

func (af *AdRequestForm) Create() {
	adRequest := new(models.AdRequest)
	adRequest.Name = af.Name
	adRequest.Mail = af.Mail
	adRequest.MessageJa = af.MessageJa
	adRequest.MessageEn = af.MessageEn
	adRequest.URLJa = af.URLJa
	adRequest.URLEn = af.URLEn
	startDate, err := time.Parse("2006-01-02", af.StartDate)
	if err != nil {
		return 
	}
	adRequest.StartDate = startDate
	endDate, err := time.Parse("2006-01-02", af.EndDate)
	if err != nil {
		return 
	}
	adRequest.EndDate = endDate
	fmt.Println(adRequest)
}

func (af *AdRequestForm) Validate() []string {
	var errorMessages []string

	validate := validator.New()

	validate.RegisterValidation("mail", mailValidation)
	validate.RegisterValidation("date", DateValidation)

	err := validate.Struct(af);

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errorMessage string

			switch err.Tag() {
			case "required":
				errorMessage = fmt.Sprintf("%s は必須です。", err.Field())
			case "mail":
				errorMessage = "メールアドレスの形式が違います"
			case "url":
				errorMessage = "url形式が違います"
			case "date":
				errorMessage = "日付の形式が違います"
			}
			errorMessages = append(errorMessages, errorMessage)
		}
	}else {
		if !PeriodValidation(af.StartDate, af.EndDate) {
			errorMessage := "期間が正しくありません"
			errorMessages = append(errorMessages, errorMessage)

		}
	}
	
	return errorMessages
}