package models

import "gorm.io/gorm"

type Medication struct {
	gorm.Model
	Name          string         `gorm:"not null;uniqueIndex"`
	Manufacturer  string         `gorm:"not null"`
	Prescriptions []Prescription `gorm:"foreignKey:MedicationID"`
}
