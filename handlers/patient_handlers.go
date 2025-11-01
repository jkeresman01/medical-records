package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/mapper"
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/repository/factory"
)

func GetPatients(c *fiber.Ctx) error {
	patientsRepository := repositoryfactory.GetInstance[models.Patient]()

	patients, err := patientsRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching patients")
	}

	patientVMs := mapper.ToPatientViewModelList(patients)

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

	patientVMs := mapper.ToPatientViewModelList(patients)

	return c.Render("patients/patient_list", fiber.Map{
		"Patients": patientVMs,
	})
}

func CreatePatient(c *fiber.Ctx) error {
	patientRepository := repositoryfactory.GetInstance[models.Patient]()

	patient := &models.Patient{
		FirstName: c.FormValue("first_name"),
		LastName:  c.FormValue("last_name"),
		DOB:       c.FormValue("dob"),
	}

	if err := patientRepository.Create(patient); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating patient")
	}

	patients, err := patientRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch patients!")
	}

	patientVMs := mapper.ToPatientViewModelList(patients)

	return c.Render("patients/patient_list", fiber.Map{
		"Patients": patientVMs,
	})
}

func GetPatient(c *fiber.Ctx) error {
	patientsRepository := repositoryfactory.GetInstance[models.Patient]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	patient, err := patientsRepository.FindByIDWithPreloads(uint(id), "Prescriptions.Medication", "Exams.ExamType")
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Patient not found")
	}

	patientDetailVM := mapper.ToPatientDetailViewModel(*patient)

	return c.Render("patients/patient_detail", fiber.Map{
		"Patient": patientDetailVM,
	})
}

func UpdatePatient(c *fiber.Ctx) error {
	patientRepository := repositoryfactory.GetInstance[models.Patient]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	patient, err := patientRepository.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Patient not found")
	}

	patient.FirstName = c.FormValue("first_name")
	patient.LastName = c.FormValue("last_name")
	patient.DOB = c.FormValue("dob")

	if err := patientRepository.Update(patient); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating patient")
	}

	patients, err := patientRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Patient not found")
	}

	patientVMs := mapper.ToPatientViewModelList(patients)

	return c.Render("patients/patient_list", fiber.Map{
		"Patients": patientVMs,
	})
}

func DeletePatient(c *fiber.Ctx) error {
	patientsRepository := repositoryfactory.GetInstance[models.Patient]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	if err := patientsRepository.DeleteByID(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting patient")
	}

	patients, err := patientsRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch patients!")
	}

	patientVMs := mapper.ToPatientViewModelList(patients)

	return c.Render("patients/patient_list", fiber.Map{
		"Patients": patientVMs,
	})
}

func GetPatientForm(c *fiber.Ctx) error {
	return c.Render("patients/patient_form", fiber.Map{})
}

func GetEditPatientForm(c *fiber.Ctx) error {
	patientsRepository := repositoryfactory.GetInstance[models.Patient]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	patient, err := patientsRepository.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Patient not found")
	}

	patientVM := mapper.ToPatientViewModel(*patient)

	return c.Render("patients/patient_edit_form", fiber.Map{
		"Patient": patientVM,
	})
}
