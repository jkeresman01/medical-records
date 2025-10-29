package models

type ExamType struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"not null;uniqueIndex"`
	Description string `gorm:"not null"`
}
