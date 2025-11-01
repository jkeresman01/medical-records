package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	FirstName     string         `gorm:"not null"`
	LastName      string         `gorm:"not null"`
	DOB           string         `gorm:"not null"`
	Prescriptions []Prescription `gorm:"foreignKey:PatientID"`
	Exams         []Exam         `gorm:"foreignKey:PatientID"`
}
