package models

type Medication struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"not null;uniqueIndex"`
	Manufacturer string `gorm:"not null"`
}
