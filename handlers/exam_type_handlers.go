package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/repository/factory"
	"github.com/jkeresman01/medical-records/viewmodels"
)

func GetExamTypes(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.ExamType]()
	examTypes, err := repo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching exam types")
	}

	var examTypeVMs []viewmodels.ExamTypeViewModel
	for _, et := range examTypes {
		examTypeVMs = append(examTypeVMs, viewmodels.ExamTypeViewModel{
			ID:          et.ID,
			Name:        et.Name,
			Description: et.Description,
		})
	}

	return c.Render("exam_types/exam_types", fiber.Map{
		"ExamTypes": examTypeVMs,
	}, "layout")
}

func GetExamTypeForm(c *fiber.Ctx) error {
	return c.Render("exam_types/exam_type_form", fiber.Map{})
}

func GetEditExamTypeForm(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.ExamType]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	examType, err := repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Exam type not found")
	}

	examTypeVM := viewmodels.ExamTypeViewModel{
		ID:          examType.ID,
		Name:        examType.Name,
		Description: examType.Description,
	}

	return c.Render("exam_types/exam_type_edit_form", fiber.Map{
		"ExamType": examTypeVM,
	})
}

func CreateExamType(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.ExamType]()

	examType := &models.ExamType{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
	}

	if err := repo.Create(examType); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating exam type")
	}

	examTypes, _ := repo.FindAll()
	var examTypeVMs []viewmodels.ExamTypeViewModel
	for _, et := range examTypes {
		examTypeVMs = append(examTypeVMs, viewmodels.ExamTypeViewModel{
			ID:          et.ID,
			Name:        et.Name,
			Description: et.Description,
		})
	}

	return c.Render("exam_types/exam_type_list", fiber.Map{
		"ExamTypes": examTypeVMs,
	})
}

func UpdateExamType(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.ExamType]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	examType, err := repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Exam type not found")
	}

	examType.Name = c.FormValue("name")
	examType.Description = c.FormValue("description")

	if err := repo.Update(examType); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating exam type")
	}

	examTypes, _ := repo.FindAll()
	var examTypeVMs []viewmodels.ExamTypeViewModel
	for _, et := range examTypes {
		examTypeVMs = append(examTypeVMs, viewmodels.ExamTypeViewModel{
			ID:          et.ID,
			Name:        et.Name,
			Description: et.Description,
		})
	}

	return c.Render("exam_types/exam_type_list", fiber.Map{
		"ExamTypes": examTypeVMs,
	})
}

func DeleteExamType(c *fiber.Ctx) error {
	repo := repositoryfactory.GetInstance[models.ExamType]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	if err := repo.DeleteByID(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting exam type")
	}

	examTypes, _ := repo.FindAll()
	var examTypeVMs []viewmodels.ExamTypeViewModel
	for _, et := range examTypes {
		examTypeVMs = append(examTypeVMs, viewmodels.ExamTypeViewModel{
			ID:          et.ID,
			Name:        et.Name,
			Description: et.Description,
		})
	}

	return c.Render("exam_types/exam_type_list", fiber.Map{
		"ExamTypes": examTypeVMs,
	})
}
