package models

import "time"

type Exam struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	PatientID  uint      `gorm:"not null;index"`
	ExamTypeID uint      `gorm:"not null;index"`
	Result     string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
