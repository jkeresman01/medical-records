package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/repository/factory"
	"github.com/jkeresman01/medical-records/viewmodels"
)

func GetMedications(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Medication]()
	medications, err := repo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching medications")
	}

	var medicationVMs []viewmodels.MedicationViewModel
	for _, m := range medications {
		medicationVMs = append(medicationVMs, viewmodels.MedicationViewModel{
			ID:           m.ID,
			Name:         m.Name,
			Manufacturer: m.Manufacturer,
		})
	}

	return c.Render("medications/medications", fiber.Map{
		"Medications": medicationVMs,
	}, "layout")
}

func GetMedicationForm(c *fiber.Ctx) error {
	return c.Render("medications/medication_form", fiber.Map{})
}

func GetEditMedicationForm(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Medication]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	medication, err := repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Medication not found")
	}

	medicationVM := viewmodels.MedicationViewModel{
		ID:           medication.ID,
		Name:         medication.Name,
		Manufacturer: medication.Manufacturer,
	}

	return c.Render("medications/medication_edit_form", fiber.Map{
		"Medication": medicationVM,
	})
}

func CreateMedication(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Medication]()

	medication := &models.Medication{
		Name:         c.FormValue("name"),
		Manufacturer: c.FormValue("manufacturer"),
	}

	if err := repo.Create(medication); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating medication")
	}

	medications, _ := repo.FindAll()
	var medicationVMs []viewmodels.MedicationViewModel
	for _, m := range medications {
		medicationVMs = append(medicationVMs, viewmodels.MedicationViewModel{
			ID:           m.ID,
			Name:         m.Name,
			Manufacturer: m.Manufacturer,
		})
	}

	return c.Render("medications/medication_list", fiber.Map{
		"Medications": medicationVMs,
	})
}

func UpdateMedication(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Medication]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	medication, err := repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Medication not found")
	}

	medication.Name = c.FormValue("name")
	medication.Manufacturer = c.FormValue("manufacturer")

	if err := repo.Update(medication); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating medication")
	}

	medications, _ := repo.FindAll()
	var medicationVMs []viewmodels.MedicationViewModel
	for _, m := range medications {
		medicationVMs = append(medicationVMs, viewmodels.MedicationViewModel{
			ID:           m.ID,
			Name:         m.Name,
			Manufacturer: m.Manufacturer,
		})
	}

	return c.Render("medications/medication_list", fiber.Map{
		"Medications": medicationVMs,
	})
}

func DeleteMedication(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.Medication]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	if err := repo.DeleteByID(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting medication")
	}

	medications, _ := repo.FindAll()
	var medicationVMs []viewmodels.MedicationViewModel
	for _, m := range medications {
		medicationVMs = append(medicationVMs, viewmodels.MedicationViewModel{
			ID:           m.ID,
			Name:         m.Name,
			Manufacturer: m.Manufacturer,
		})
	}

	return c.Render("medications/medication_list", fiber.Map{
		"Medications": medicationVMs,
	})
}
