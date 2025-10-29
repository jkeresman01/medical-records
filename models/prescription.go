package models

type Prescription struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	PatientID    uint   `gorm:"not null;index"`
	MedicationID uint   `gorm:"not null;index"`
	Dosage       string `gorm:"not null"`
	Frequency    string `gorm:"not null"`
}
