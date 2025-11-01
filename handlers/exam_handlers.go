package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/mapper"
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/repository/factory"
)

func GetExams(c *fiber.Ctx) error {
	examRepository := repositoryfactory.GetInstance[models.Exam]()
	patientRepository := repositoryfactory.GetInstance[models.Patient]()
	examTypeRepository := repositoryfactory.GetInstance[models.ExamType]()

	exams, err := examRepository.FindAllWithPreloads("Patient", "ExamType")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching exams")
	}

	examVMs := mapper.ToExamViewModelList(exams)

	patients, err := patientRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed fetching patients")
	}

	patientVMs := mapper.ToPatientViewModelList(patients)

	examTypes, err := examTypeRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Failed exam types patients")
	}

	examTypeVMs := mapper.ToExamTypeViewModelList(examTypes)

	return c.Render("exams/exams", fiber.Map{
		"Exams":     examVMs,
		"Patients":  patientVMs,
		"ExamTypes": examTypeVMs,
	}, "layout")
}

func GetExamForm(c *fiber.Ctx) error {
	patientRepository := repositoryfactory.GetInstance[models.Patient]()
	examTypeRepository := repositoryfactory.GetInstance[models.ExamType]()

	patients, err := patientRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching exams")
	}

	patientVMs := mapper.ToPatientViewModelList(patients)

	examTypes, err := examTypeRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Failed fetching exam types")
	}

	examTypeVMs := mapper.ToExamTypeViewModelList(examTypes)

	return c.Render("exams/exam_form", fiber.Map{
		"Patients":  patientVMs,
		"ExamTypes": examTypeVMs,
	})
}

func GetEditExamForm(c *fiber.Ctx) error {
	examRepository := repositoryfactory.GetInstance[models.Exam]()
	patientRepository := repositoryfactory.GetInstance[models.Patient]()
	examTypeRepository := repositoryfactory.GetInstance[models.ExamType]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	exam, err := examRepository.FindByIDWithPreloads(uint(id), "Patient", "ExamType")
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Exam not found")
	}

	examVM := mapper.ToExamViewModel(*exam)

	patients, err := patientRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed fetching patients")
	}

	patientVMs := mapper.ToPatientViewModelList(patients)

	examTypes, err := examTypeRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Failed exam types patients")
	}

	examTypeVMs := mapper.ToExamTypeViewModelList(examTypes)

	return c.Render("exams/exam_edit_form", fiber.Map{
		"Exam":          examVM,
		"ExamPatientID": exam.Patient.ID,
		"ExamTypeID":    exam.ExamType.ID,
		"Patients":      patientVMs,
		"ExamTypes":     examTypeVMs,
	})
}

func CreateExam(c *fiber.Ctx) error {
	examRepository := repositoryfactory.GetInstance[models.Exam]()

	patientID, _ := strconv.ParseUint(c.FormValue("patient_id"), 10, 32)
	examTypeID, _ := strconv.ParseUint(c.FormValue("exam_type_id"), 10, 32)

	exam := &models.Exam{
		PatientID:  uint(patientID),
		ExamTypeID: uint(examTypeID),
		Result:     c.FormValue("result"),
	}

	if err := examRepository.Create(exam); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating exam")
	}

	exams, _ := examRepository.FindAllWithPreloads("Patient", "ExamType")
	examVMs := mapper.ToExamViewModelList(exams)

	return c.Render("exams/exam_list", fiber.Map{
		"Exams": examVMs,
	})
}

func UpdateExam(c *fiber.Ctx) error {
	examRepository := repositoryfactory.GetInstance[models.Exam]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	exam, err := examRepository.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Exam not found")
	}

	patientID, _ := strconv.ParseUint(c.FormValue("patient_id"), 10, 32)
	examTypeID, _ := strconv.ParseUint(c.FormValue("exam_type_id"), 10, 32)

	exam.PatientID = uint(patientID)
	exam.ExamTypeID = uint(examTypeID)
	exam.Result = c.FormValue("result")

	if err := examRepository.Update(exam); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating exam")
	}

	exams, err := examRepository.FindAllWithPreloads("Patient", "ExamType")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch exams!")
	}

	examVMs := mapper.ToExamViewModelList(exams)

	return c.Render("exams/exam_list", fiber.Map{
		"Exams": examVMs,
	})
}

func DeleteExam(c *fiber.Ctx) error {
	examRepository := repositoryfactory.GetInstance[models.Exam]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	if err := examRepository.DeleteByID(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting exam")
	}

	exams, err := examRepository.FindAllWithPreloads("Patient", "ExamType")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch exams!")
	}

	examVMs := mapper.ToExamViewModelList(exams)

	return c.Render("exams/exam_list", fiber.Map{
		"Exams": examVMs,
	})
}
