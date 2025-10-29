package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jkeresman01/medical-records/config"
	"github.com/jkeresman01/medical-records/models"
)

var DB                *gorm.DB

func Connect(cfg *config.Config) {
	db, err := gorm.Open(postgres.Open(cfg.ConnString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.Patient{},
		&models.Medication{},
		&models.Prescription{},
		&models.ExamType{},
		&models.Exam{},
	); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	DB = db
}
