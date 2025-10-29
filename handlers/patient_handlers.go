package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/db"
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/repository/factory"
	"github.com/jkeresman01/medical-records/viewmodels"
)

func GetPatients(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Patient]()
	patients, err := repo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching patients")
	}

	var patientVMs []viewmodels.PatientViewModel
	for _, p := range patients {
		patientVMs = append(patientVMs, viewmodels.PatientViewModel{
			ID:        p.ID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			DOB:       p.DOB,
		})
	}

	return c.Render("patients/patients", fiber.Map{
		"Patients": patientVMs,
	}, "layout")
}

func GetPatientsList(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Patient]()
	patients, err := repo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching patients")
	}

	var patientVMs []viewmodels.PatientViewModel
	for _, p := range patients {
		patientVMs = append(patientVMs, viewmodels.PatientViewModel{
			ID:        p.ID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			DOB:       p.DOB,
		})
	}

	return c.Render("patients/patient_list", fiber.Map{
		"Patients": patientVMs,
	})
}

func CreatePatient(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Patient]()

	patient := &models.Patient{
		FirstName: c.FormValue("first_name"),
		LastName:  c.FormValue("last_name"),
		DOB:       c.FormValue("dob"),
	}

	if err := repo.Create(patient); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating patient")
	}

	patients, _ := repo.FindAll()
	var patientVMs []viewmodels.PatientViewModel
	for _, p := range patients {
		patientVMs = append(patientVMs, viewmodels.PatientViewModel{
			ID:        p.ID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			DOB:       p.DOB,
		})
	}

	return c.Render("patients/patient_list", fiber.Map{
		"Patients": patientVMs,
	})
}

func GetPatient(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Patient]()
	medicationRepo := repositoryfactory.GetInstance[models.Medication]()
	examTypeRepo := repositoryfactory.GetInstance[models.ExamType]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	patient, err := repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Patient not found")
	}

	var allPrescriptions []models.Prescription
	db.DB.Where("patient_id = ?", id).Find(&allPrescriptions)

	var prescriptionVMs []viewmodels.PrescriptionViewModel
	for _, p := range allPrescriptions {
		medication, _ := medicationRepo.FindByID(p.MedicationID)
		prescriptionVMs = append(prescriptionVMs, viewmodels.PrescriptionViewModel{
			ID:             p.ID,
			PatientID:      p.PatientID,
			MedicationID:   p.MedicationID,
			PatientName:    patient.FirstName + " " + patient.LastName,
			MedicationName: medication.Name,
			Dosage:         p.Dosage,
			Frequency:      p.Frequency,
		})
	}

	var allExams []models.Exam
	db.DB.Where("patient_id = ?", id).Find(&allExams)

	var examVMs []viewmodels.ExamViewModel
	for _, e := range allExams {
		examType, _ := examTypeRepo.FindByID(e.ExamTypeID)
		examVMs = append(examVMs, viewmodels.ExamViewModel{
			ID:           e.ID,
			PatientID:    e.PatientID,
			ExamTypeID:   e.ExamTypeID,
			PatientName:  patient.FirstName + " " + patient.LastName,
			ExamTypeName: examType.Name,
			Result:       e.Result,
			CreatedAt:    e.CreatedAt.Format("2006-01-02"),
		})
	}

	patientDetailVM := viewmodels.PatientDetailViewModel{
		ID:            patient.ID,
		FirstName:     patient.FirstName,
		LastName:      patient.LastName,
		DOB:           patient.DOB,
		Prescriptions: prescriptionVMs,
		Exams:         examVMs,
	}

	return c.Render("patients/patient_detail", fiber.Map{
		"Patient": patientDetailVM,
	})
}

func UpdatePatient(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Patient]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("No can do, Invalid ID")
	}

	patient, err := repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("No can do, Patient not found")
	}

	patient.FirstName = c.FormValue("first_name")
	patient.LastName = c.FormValue("last_name")
	patient.DOB = c.FormValue("dob")

	if err := repo.Update(patient); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating patient")
	}

	patients, _ := repo.FindAll()
	var patientVMs []viewmodels.PatientViewModel
	for _, p := range patients {
		patientVMs = append(patientVMs, viewmodels.PatientViewModel{
			ID:        p.ID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			DOB:       p.DOB,
		})
	}

	return c.Render("patients/patient_list", fiber.Map{
		"Patients": patientVMs,
	})
}

func DeletePatient(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Patient]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	if err := repo.DeleteByID(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting patient")
	}

	patients, _ := repo.FindAll()
	var patientVMs []viewmodels.PatientViewModel
	for _, p := range patients {
		patientVMs = append(patientVMs, viewmodels.PatientViewModel{
			ID:        p.ID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			DOB:       p.DOB,
		})
	}

	return c.Render("patients/patient_list", fiber.Map{
		"Patients": patientVMs,
	})
}

func GetPatientForm(c *fiber.Ctx) error {
	return c.Render("patients/patient_form", fiber.Map{})
}

func GetEditPatientForm(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Patient]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("No can do, Invalid ID")
	}

	patient, err := repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Patient not found")
	}

	patientVM := viewmodels.PatientViewModel{
		ID:        patient.ID,
		FirstName: patient.FirstName,
		LastName:  patient.LastName,
		DOB:       patient.DOB,
	}

	return c.Render("patients/patient_edit_form", fiber.Map{
		"Patient": patientVM,
	})
}
