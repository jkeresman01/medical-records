package models

import "gorm.io/gorm"

type Prescription struct {
	gorm.Model
	PatientID    uint
	Patient      Patient
	MedicationID uint
	Medication   Medication
	Dosage       string `gorm:"not null"`
	Frequency    string `gorm:"not null"`
}
