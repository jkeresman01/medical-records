package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/mapper"
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/repository/factory"
)

func GetPrescriptions(c *fiber.Ctx) error {
	prescriptionsRepository := repositoryfactory.GetInstance[models.Prescription]()
	patientRepository := repositoryfactory.GetInstance[models.Patient]()
	medicationRepository := repositoryfactory.GetInstance[models.Medication]()

	prescriptions, err := prescriptionsRepository.FindAllWithPreloads("Patient", "Medication")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching prescriptions")
	}

	prescriptionVMs := mapper.ToPrescriptionViewModelList(prescriptions)

	patients, err := patientRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch patietns!")
	}

	patientVMs := mapper.ToPatientViewModelList(patients)

	medications, err := medicationRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch medications!")
	}

	medicationVMs := mapper.ToMedicationViewModelList(medications)

	return c.Render("prescriptions/prescriptions", fiber.Map{
		"Prescriptions": prescriptionVMs,
		"Patients":      patientVMs,
		"Medications":   medicationVMs,
	}, "layout")
}

func GetPrescriptionForm(c *fiber.Ctx) error {
	patientRepository := repositoryfactory.GetInstance[models.Patient]()
	medicationRepository := repositoryfactory.GetInstance[models.Medication]()

	patients, err := patientRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch patietns!")
	}

	patientVMs := mapper.ToPatientViewModelList(patients)

	medications, err := medicationRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch medications!")
	}

	medicationVMs := mapper.ToMedicationViewModelList(medications)

	return c.Render("prescriptions/prescription_form", fiber.Map{
		"Patients":    patientVMs,
		"Medications": medicationVMs,
	})
}

func GetEditPrescriptionForm(c *fiber.Ctx) error {
	prescriptionRepository := repositoryfactory.GetInstance[models.Prescription]()
	patientRepository := repositoryfactory.GetInstance[models.Patient]()
	medicationsRepository := repositoryfactory.GetInstance[models.Medication]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	prescription, err := prescriptionRepository.FindByIDWithPreloads(uint(id), "Patient", "Medication")
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Prescription not found")
	}

	prescriptionVM := mapper.ToPrescriptionViewModel(*prescription)

	patients, err := patientRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch patietns!")
	}

	patientVMs := mapper.ToPatientViewModelList(patients)

	medications, err := medicationsRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch medications!")
	}

	medicationVMs := mapper.ToMedicationViewModelList(medications)

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

	patientID, err := strconv.ParseUint(c.FormValue("patient_id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	medicationID, err := strconv.ParseUint(c.FormValue("medication_id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

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

	prescriptions, err := prescriptionRepository.FindAllWithPreloads("Patient", "Medication")
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Prescription not found")
	}

	prescriptionVMs := mapper.ToPrescriptionViewModelList(prescriptions)

	return c.Render("prescriptions/prescription_list", fiber.Map{
		"Prescriptions": prescriptionVMs,
	})
}

func UpdatePrescription(c *fiber.Ctx) error {
	prescriptionRepository := repositoryfactory.GetInstance[models.Prescription]()

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

	prescription.PatientID = uint(patientID)
	prescription.MedicationID = uint(medicationID)
	prescription.Dosage = c.FormValue("dosage")
	prescription.Frequency = c.FormValue("frequency")

	if err := prescriptionRepository.Update(prescription); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating prescription")
	}

	prescriptions, err := prescriptionRepository.FindAllWithPreloads("Patient", "Medication")
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Prescription not found")
	}

	prescriptionVMs := mapper.ToPrescriptionViewModelList(prescriptions)

	return c.Render("prescriptions/prescription_list", fiber.Map{
		"Prescriptions": prescriptionVMs,
	})
}

func DeletePrescription(c *fiber.Ctx) error {
	prescriptionRepository := repositoryfactory.GetInstance[models.Prescription]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	if err := prescriptionRepository.DeleteByID(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting prescription")
	}

	prescriptions, err := prescriptionRepository.FindAllWithPreloads("Patient", "Medication")
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Prescription not found")
	}

	prescriptionVMs := mapper.ToPrescriptionViewModelList(prescriptions)

	return c.Render("prescriptions/prescription_list", fiber.Map{
		"Prescriptions": prescriptionVMs,
	})
}
