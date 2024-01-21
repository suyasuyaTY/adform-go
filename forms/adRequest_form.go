package forms

import (
	"fmt"

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
	fmt.Println(af.Name)
	fmt.Println(af.Mail)
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
	}
	
	return errorMessages
}