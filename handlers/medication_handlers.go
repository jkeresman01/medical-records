package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/mapper"
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/repository/factory"
)

func GetMedications(c *fiber.Ctx) error {
	medicationsRepository := repositoryfactory.GetInstance[models.Medication]()

	medications, err := medicationsRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching medications")
	}

	medicationVMs := mapper.ToMedicationViewModelList(medications)

	return c.Render("medications/medications", fiber.Map{
		"Medications": medicationVMs,
	}, "layout")
}

func GetMedicationForm(c *fiber.Ctx) error {
	return c.Render("medications/medication_form", fiber.Map{})
}

func GetEditMedicationForm(c *fiber.Ctx) error {
	medicationRepository := repositoryfactory.GetInstance[models.Medication]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	medication, err := medicationRepository.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Medication not found")
	}

	medicationVM := mapper.ToMedicationViewModel(*medication)

	return c.Render("medications/medication_edit_form", fiber.Map{
		"Medication": medicationVM,
	})
}

func CreateMedication(c *fiber.Ctx) error {
	medicationsRepository := repositoryfactory.GetInstance[models.Medication]()

	medication := &models.Medication{
		Name:         c.FormValue("name"),
		Manufacturer: c.FormValue("manufacturer"),
	}

	if err := medicationsRepository.Create(medication); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating medication")
	}

	medications, err := medicationsRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch all medications!")
	}

	medicationVMs := mapper.ToMedicationViewModelList(medications)

	return c.Render("medications/medication_list", fiber.Map{
		"Medications": medicationVMs,
	})
}

func UpdateMedication(c *fiber.Ctx) error {
	medicationRepository := repositoryfactory.GetInstance[models.Medication]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	medication, err := medicationRepository.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Medication not found")
	}

	medication.Name = c.FormValue("name")
	medication.Manufacturer = c.FormValue("manufacturer")

	if err := medicationRepository.Update(medication); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating medication")
	}

	medications, err := medicationRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch medication!")
	}

	medicationVMs := mapper.ToMedicationViewModelList(medications)

	return c.Render("medications/medication_list", fiber.Map{
		"Medications": medicationVMs,
	})
}

func DeleteMedication(c *fiber.Ctx) error {
	medicationsRepository := repositoryfactory.GetInstance[models.Medication]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	if err := medicationsRepository.DeleteByID(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting medication")
	}

	medications, err := medicationsRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch medications!")
	}

	medicationVMs := mapper.ToMedicationViewModelList(medications)

	return c.Render("medications/medication_list", fiber.Map{
		"Medications": medicationVMs,
	})
}
