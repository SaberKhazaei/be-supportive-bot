package model

import (
	"time"
)

type BaleBot struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	Stat         string
	Number       int64
	NationalCode int64
	BirthDate    time.Time
}
