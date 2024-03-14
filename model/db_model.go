package model

import (
	"time"
)

type BaleBot struct {
	ID                uint `json:"id" gorm:"primaryKey"`
	Stat              string
	PhoneNumber       int64
	FirstName         string
	LastName          string
	NationalCode      int64
	JobTitle          string
	VerificationToken string
	VerificationCode  int64
	BirthDate         time.Time
}
