package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/repository/factory"
	"github.com/jkeresman01/medical-records/viewmodels"
)

func GetPrescriptions(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Prescription]()
	patientRepo := repositoryfactory.GetInstance[models.Patient]()
	medicationRepo := repositoryfactory.GetInstance[models.Medication]()

	prescriptions, err := repo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching prescriptions")
	}

	var prescriptionVMs []viewmodels.PrescriptionViewModel
	for _, p := range prescriptions {
		patient, _ := patientRepo.FindByID(p.PatientID)
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

	patients, _ := patientRepo.FindAll()
	var patientVMs []viewmodels.PatientViewModel
	for _, pat := range patients {
		patientVMs = append(patientVMs, viewmodels.PatientViewModel{
			ID:        pat.ID,
			FirstName: pat.FirstName,
			LastName:  pat.LastName,
			DOB:       pat.DOB,
		})
	}

	medications, _ := medicationRepo.FindAll()
	var medicationVMs []viewmodels.MedicationViewModel
	for _, med := range medications {
		medicationVMs = append(medicationVMs, viewmodels.MedicationViewModel{
			ID:           med.ID,
			Name:         med.Name,
			Manufacturer: med.Manufacturer,
		})
	}

	return c.Render("prescriptions/prescriptions", fiber.Map{
		"Prescriptions": prescriptionVMs,
		"Patients":      patientVMs,
		"Medications":   medicationVMs,
	}, "layout")
}

func GetPrescriptionForm(c *fiber.Ctx) error {
	patientRepo := repositoryfactory.GetInstance[models.Patient]()
	medicationRepo := repositoryfactory.GetInstance[models.Medication]()

	patients, _ := patientRepo.FindAll()
	var patientVMs []viewmodels.PatientViewModel
	for _, p := range patients {
		patientVMs = append(patientVMs, viewmodels.PatientViewModel{
			ID:        p.ID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			DOB:       p.DOB,
		})
	}

	medications, _ := medicationRepo.FindAll()
	var medicationVMs []viewmodels.MedicationViewModel
	for _, m := range medications {
		medicationVMs = append(medicationVMs, viewmodels.MedicationViewModel{
			ID:           m.ID,
			Name:         m.Name,
			Manufacturer: m.Manufacturer,
		})
	}

	return c.Render("prescriptions/prescription_form", fiber.Map{
		"Patients":    patientVMs,
		"Medications": medicationVMs,
	})
}

func GetEditPrescriptionForm(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Prescription]()
	patientRepo := repositoryfactory.GetInstance[models.Patient]()
	medicationRepo := repositoryfactory.GetInstance[models.Medication]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	prescription, err := repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Prescription not found")
	}

	patient, _ := patientRepo.FindByID(prescription.PatientID)
	medication, _ := medicationRepo.FindByID(prescription.MedicationID)

	prescriptionVM := viewmodels.PrescriptionViewModel{
		ID:             prescription.ID,
		PatientID:      prescription.PatientID,
		MedicationID:   prescription.MedicationID,
		PatientName:    patient.FirstName + " " + patient.LastName,
		MedicationName: medication.Name,
		Dosage:         prescription.Dosage,
		Frequency:      prescription.Frequency,
	}

	patients, _ := patientRepo.FindAll()
	var patientVMs []viewmodels.PatientViewModel
	for _, p := range patients {
		patientVMs = append(patientVMs, viewmodels.PatientViewModel{
			ID:        p.ID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			DOB:       p.DOB,
		})
	}

	medications, _ := medicationRepo.FindAll()
	var medicationVMs []viewmodels.MedicationViewModel
	for _, m := range medications {
		medicationVMs = append(medicationVMs, viewmodels.MedicationViewModel{
			ID:           m.ID,
			Name:         m.Name,
			Manufacturer: m.Manufacturer,
		})
	}

	return c.Render("prescriptions/prescription_edit_form", fiber.Map{
		"Prescription": prescriptionVM,
		"Patients":     patientVMs,
		"Medications":  medicationVMs,
	})
}

func CreatePrescription(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Prescription]()
	patientRepo := repositoryfactory.GetInstance[models.Patient]()
	medicationRepo := repositoryfactory.GetInstance[models.Medication]()

	patientID, _ := strconv.ParseUint(c.FormValue("patient_id"), 10, 32)
	medicationID, _ := strconv.ParseUint(c.FormValue("medication_id"), 10, 32)

	prescription := &models.Prescription{
		PatientID:    uint(patientID),
		MedicationID: uint(medicationID),
		Dosage:       c.FormValue("dosage"),
		Frequency:    c.FormValue("frequency"),
	}

	if err := repo.Create(prescription); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating prescription")
	}

	prescriptions, _ := repo.FindAll()
	var prescriptionVMs []viewmodels.PrescriptionViewModel
	for _, p := range prescriptions {
		patient, _ := patientRepo.FindByID(p.PatientID)
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

	return c.Render("prescriptions/prescription_list", fiber.Map{
		"Prescriptions": prescriptionVMs,
	})
}

func UpdatePrescription(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Prescription]()
	patientRepo := repositoryfactory.GetInstance[models.Patient]()
	medicationRepo := repositoryfactory.GetInstance[models.Medication]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("No can do, Invalid ID")
	}

	prescription, err := repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Prescription not found")
	}

	patientID, _ := strconv.ParseUint(c.FormValue("patient_id"), 10, 32)
	medicationID, _ := strconv.ParseUint(c.FormValue("medication_id"), 10, 32)

	prescription.PatientID = uint(patientID)
	prescription.MedicationID = uint(medicationID)
	prescription.Dosage = c.FormValue("dosage")
	prescription.Frequency = c.FormValue("frequency")

	if err := repo.Update(prescription); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating prescription")
	}

	prescriptions, _ := repo.FindAll()
	var prescriptionVMs []viewmodels.PrescriptionViewModel
	for _, p := range prescriptions {
		patient, _ := patientRepo.FindByID(p.PatientID)
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

	return c.Render("prescriptions/prescription_list", fiber.Map{
		"Prescriptions": prescriptionVMs,
	})
}

func DeletePrescription(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Prescription]()
	patientRepo := repositoryfactory.GetInstance[models.Patient]()
	medicationRepo := repositoryfactory.GetInstance[models.Medication]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	if err := repo.DeleteByID(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting prescription")
	}

	prescriptions, _ := repo.FindAll()
	var prescriptionVMs []viewmodels.PrescriptionViewModel
	for _, p := range prescriptions {
		patient, _ := patientRepo.FindByID(p.PatientID)
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

	return c.Render("prescriptions/prescription_list", fiber.Map{
		"Prescriptions": prescriptionVMs,
	})
}

