package forms

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

func mailValidation(fl validator.FieldLevel) bool {
	mMailPattern := `(^[a-z]+\.[a-z]\.[a-z]{2}@m\.titech\.ac\.jp$)`

	re, err := regexp.Compile(mMailPattern)
	if err != nil {
		return false
	}
	return re.MatchString(fl.Field().String())
}

func DateValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	return err == nil
}