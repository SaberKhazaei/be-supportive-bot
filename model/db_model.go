package model

import (
	"gorm.io/datatypes"
	"time"
)

type BaleBot struct {
	ID                uint `json:"id" gorm:"primaryKey"`
	FirstName         string
	LastName          string
	PhoneNumber       string
	NationalCode      string
	Password          string
	EnteredTime       *time.Time
	Stat              string
	Captcha           string
	JobId             string
	VerificationToken string
	VerificationCode  string
	BirthDate         string
	SiteCookie        string
	JobIdLoginCode    string
	RepresentedChild  datatypes.JSON
	CurrentChildInfo  datatypes.JSON
}
