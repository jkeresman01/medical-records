package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jkeresman01/medical-records/config"
	"github.com/jkeresman01/medical-records/models"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	db, err := gorm.Open(postgres.Open(cfg.ConnString()), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
}

func Migrate() {
	if err := DB.AutoMigrate(
		&models.Patient{},
		&models.Medication{},
		&models.Prescription{},
		&models.ExamType{},
		&models.Exam{},
	); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	seedExamTypes(DB)
}

func seedExamTypes(db *gorm.DB) {
	for _, examType := range models.PredefinedExamTypes {
		var existing models.ExamType
		result := db.Where("name = ?", examType.Name).First(&existing)

		if result.Error == gorm.ErrRecordNotFound {
			if err := db.Create(&examType).Error; err != nil {
				log.Printf("Warning: Failed to create exam type %s: %v", examType.Name, err)
			}
		}
	}
	log.Println("Exam types seeded successfully")
}
