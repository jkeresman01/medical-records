package models

import (
	"gorm.io/gorm"
	"time"
)

type Exam struct {
	gorm.Model
	PatientID  uint
	Patient    Patient
	ExamTypeID uint
	ExamType   ExamType
	Result     string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
