package forms

import (
	"fmt"
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

func PeriodValidation(start string, end string) bool {
	startDate, err := time.Parse("2006-01-02", start)
	if err != nil {
		return false
	}
	endDate, err := time.Parse("2006-01-02", end)
	if err != nil {
		return false
	}
	period := endDate.Sub(startDate).Hours()/24
	fmt.Println(period)
	if period < 0 || period > 14 {
		return false
	}
	return true
}