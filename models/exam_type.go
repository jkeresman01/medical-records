package models

import "gorm.io/gorm"

type ExamType struct {
	gorm.Model
	Name        string `gorm:"not null;uniqueIndex"`
	Description string `gorm:"not null"`
	Exams       []Exam `gorm:"foreignKey:ExamTypeID"`
}

var PredefinedExamTypes = []ExamType{
	{Name: "GENERAL_PRACTITIONER", Description: "General Practitioner Examination"},
	{Name: "BLOOD_TEST", Description: "Blood Test"},
	{Name: "X_RAY", Description: "X-Ray Imaging"},
	{Name: "CT_SCAN", Description: "CT Scan"},
	{Name: "MRI", Description: "MRI Scan"},
	{Name: "ULTRASOUND", Description: "Ultrasound"},
	{Name: "ECG", Description: "Electrocardiogram"},
	{Name: "ECHOCARDIOGRAM", Description: "Echocardiogram"},
	{Name: "OPHTHALMOLOGY", Description: "Eye Examination"},
	{Name: "DERMATOLOGY", Description: "Dermatology Examination"},
	{Name: "DENTAL", Description: "Dental Examination"},
	{Name: "MAMMOGRAPHY", Description: "Mammography"},
	{Name: "EEG", Description: "Electroencephalogram"},
}
