package models

type Patient struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	DOB       string `gorm:"not null"`
}
