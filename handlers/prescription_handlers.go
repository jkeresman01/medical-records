package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/repository/factory"
	"github.com/jkeresman01/medical-records/viewmodels"
)

func GetPrescriptions(c *fiber.Ctx) error {
	prescriptionsRepository := repositoryfactory.GetInstance[models.Prescription]()
	patientRepository := repositoryfactory.GetInstance[models.Patient]()
	medicationRepository := repositoryfactory.GetInstance[models.Medication]()

	prescriptions, err := prescriptionsRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching prescriptions")
	}

	var prescriptionVMs []viewmodels.PrescriptionViewModel
	for _, p := range prescriptions {
		prescriptionVMs = append(prescriptionVMs, viewmodels.PrescriptionViewModel{
			ID:             p.ID,
			PatientName:    p.Patient.FirstName + " " + p.Patient.LastName,
			MedicationName: p.Medication.Name,
			Dosage:         p.Dosage,
			Frequency:      p.Frequency,
		})
	}

	patients, _ := patientRepository.FindAll()
	var patientVMs []viewmodels.PatientViewModel
	for _, pat := range patients {
		patientVMs = append(patientVMs, viewmodels.PatientViewModel{
			ID:        pat.ID,
			FirstName: pat.FirstName,
			LastName:  pat.LastName,
			DOB:       pat.DOB,
		})
	}

	medications, _ := medicationRepository.FindAll()
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
	prescriptionRepository := repositoryfactory.GetInstance[models.Prescription]()
	patientRepository := repositoryfactory.GetInstance[models.Patient]()
	medicationReposiotry := repositoryfactory.GetInstance[models.Medication]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	prescription, err := prescriptionRepository.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Prescription not found")
	}

	prescriptionVM := viewmodels.PrescriptionViewModel{
		ID:             prescription.ID,
		PatientName:    prescription.Patient.FirstName + " " + prescription.Patient.LastName,
		MedicationName: prescription.Medication.Name,
		Dosage:         prescription.Dosage,
		Frequency:      prescription.Frequency,
	}

	patients, _ := patientRepository.FindAll()
	var patientVMs []viewmodels.PatientViewModel
	for _, p := range patients {
		patientVMs = append(patientVMs, viewmodels.PatientViewModel{
			ID:        p.ID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			DOB:       p.DOB,
		})
	}

	medications, _ := medicationReposiotry.FindAll()
	var medicationVMs []viewmodels.MedicationViewModel
	for _, m := range medications {
		medicationVMs = append(medicationVMs, viewmodels.MedicationViewModel{
			ID:           m.ID,
			Name:         m.Name,
			Manufacturer: m.Manufacturer,
		})
	}

	return c.Render("prescriptions/prescription_edit_form", fiber.Map{
		"Prescription":             prescriptionVM,
		"PrescriptionPatientID":    prescription.Patient.ID,
		"PrescriptionMedicationID": prescription.Medication.ID,
		"Patients":                 patientVMs,
		"Medications":              medicationVMs,
	})
}

func CreatePrescription(c *fiber.Ctx) error {
	prescriptionRepository := repositoryfactory.GetInstance[models.Prescription]()
	medicationRepository := repositoryfactory.GetInstance[models.Medication]()
	patientRepository := repositoryfactory.GetInstance[models.Patient]()

	patientID, _ := strconv.ParseUint(c.FormValue("patient_id"), 10, 32)
	medicationID, _ := strconv.ParseUint(c.FormValue("medication_id"), 10, 32)

	medication, err := medicationRepository.FindByID(uint(medicationID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("No can do for medication")
	}

	patient, err := patientRepository.FindByID(uint(patientID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("No can do for patient")
	}

	prescription := &models.Prescription{
		PatientID:    patient.ID,
		MedicationID: medication.ID,
		Dosage:       c.FormValue("dosage"),
		Frequency:    c.FormValue("frequency"),
	}

	if err := prescriptionRepository.Create(prescription); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating prescription")
	}

	prescriptions, _ := prescriptionRepository.FindAll()
	var prescriptionVMs []viewmodels.PrescriptionViewModel
	for _, p := range prescriptions {
		prescriptionVMs = append(prescriptionVMs, viewmodels.PrescriptionViewModel{
			ID:             p.ID,
			PatientName:    p.Patient.FirstName + " " + p.Patient.LastName,
			MedicationName: p.Medication.Name,
			Dosage:         p.Dosage,
			Frequency:      p.Frequency,
		})
	}

	return c.Render("prescriptions/prescription_list", fiber.Map{
		"Prescriptions": prescriptionVMs,
	})
}

func UpdatePrescription(c *fiber.Ctx) error {
	prescriptionRepository := repositoryfactory.GetInstance[models.Prescription]()
	medicationRepository := repositoryfactory.GetInstance[models.Medication]()
	patientRepository := repositoryfactory.GetInstance[models.Patient]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("No can do, Invalid ID")
	}

	prescription, err := prescriptionRepository.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Prescription not found")
	}

	patientID, _ := strconv.ParseUint(c.FormValue("patient_id"), 10, 32)
	medicationID, _ := strconv.ParseUint(c.FormValue("medication_id"), 10, 32)

	patient, err := patientRepository.FindByID(uint(patientID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("No can do for patient")
	}

	medication, err := medicationRepository.FindByID(uint(medicationID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("No can do for medication")
	}

	prescription.Patient = *patient
	prescription.Medication = *medication
	prescription.Dosage = c.FormValue("dosage")
	prescription.Frequency = c.FormValue("frequency")

	if err := prescriptionRepository.Update(prescription); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating prescription")
	}

	prescriptions, _ := prescriptionRepository.FindAll()
	var prescriptionVMs []viewmodels.PrescriptionViewModel
	for _, p := range prescriptions {
		prescriptionVMs = append(prescriptionVMs, viewmodels.PrescriptionViewModel{
			ID:             p.ID,
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
	prescirptionRepository := repositoryfactory.GetInstance[models.Prescription]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	if err := prescirptionRepository.DeleteByID(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting prescription")
	}

	prescriptions, _ := prescirptionRepository.FindAll()
	var prescriptionVMs []viewmodels.PrescriptionViewModel
	for _, p := range prescriptions {
		prescriptionVMs = append(prescriptionVMs, viewmodels.PrescriptionViewModel{
			ID:             p.ID,
			PatientName:    p.Patient.FirstName + " " + p.Patient.LastName,
			MedicationName: p.Medication.Name,
			Dosage:         p.Dosage,
			Frequency:      p.Frequency,
		})
	}

	return c.Render("prescriptions/prescription_list", fiber.Map{
		"Prescriptions": prescriptionVMs,
	})
}
