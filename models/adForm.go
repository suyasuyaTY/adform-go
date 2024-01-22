package models

import "time"

type AdRequest struct {
	Name string
	Mail string
	MessageJa	string 
	MessageEn	string 
	URLJa		string 
	URLEn		string 
	StartDate	time.Time
	EndDate		time.Time
}