package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/repository/factory"
	"github.com/jkeresman01/medical-records/viewmodels"
)

func GetExams(c *fiber.Ctx) error {
	examRepository := repositoryfactory.GetInstance[models.Exam]()
	patientRepository := repositoryfactory.GetInstance[models.Patient]()
	examTypeRepository := repositoryfactory.GetInstance[models.ExamType]()

	exams, err := examRepository.FindAllWithPreloads("Patient", "ExamType")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching exams")
	}

	var examVMs []viewmodels.ExamViewModel
	for _, e := range exams {
		examVMs = append(examVMs, viewmodels.ExamViewModel{
			ID:           e.ID,
			PatientName:  e.Patient.FirstName + " " + e.Patient.LastName,
			ExamTypeName: e.ExamType.Description,
			Result:       e.Result,
			CreatedAt:    e.CreatedAt.Format("2006-01-02"),
		})
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

	examTypes, _ := examTypeRepository.FindAll()
	var examTypeVMs []viewmodels.ExamTypeViewModel
	for _, et := range examTypes {
		examTypeVMs = append(examTypeVMs, viewmodels.ExamTypeViewModel{
			ID:          et.ID,
			Name:        et.Name,
			Description: et.Description,
		})
	}

	return c.Render("exams/exams", fiber.Map{
		"Exams":     examVMs,
		"Patients":  patientVMs,
		"ExamTypes": examTypeVMs,
	}, "layout")
}

func GetExamForm(c *fiber.Ctx) error {
	patientRepository := repositoryfactory.GetInstance[models.Patient]()
	examTypeRepository := repositoryfactory.GetInstance[models.ExamType]()

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

	examTypes, _ := examTypeRepository.FindAll()
	var examTypeVMs []viewmodels.ExamTypeViewModel
	for _, et := range examTypes {
		examTypeVMs = append(examTypeVMs, viewmodels.ExamTypeViewModel{
			ID:          et.ID,
			Name:        et.Name,
			Description: et.Description,
		})
	}

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

	examVM := viewmodels.ExamViewModel{
		ID:           exam.ID,
		PatientName:  exam.Patient.FirstName + " " + exam.Patient.LastName,
		ExamTypeName: exam.ExamType.Name,
		Result:       exam.Result,
		CreatedAt:    exam.CreatedAt.Format("2006-01-02"),
	}

	patients, _ := patientRepository.FindAll()
	examTypes, _ := examTypeRepository.FindAll()

	var patientVMs []viewmodels.PatientViewModel
	for _, p := range patients {
		patientVMs = append(patientVMs, viewmodels.PatientViewModel{
			ID:        p.ID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			DOB:       p.DOB,
		})
	}

	var examTypeVMs []viewmodels.ExamTypeViewModel
	for _, et := range examTypes {
		examTypeVMs = append(examTypeVMs, viewmodels.ExamTypeViewModel{
			ID:          et.ID,
			Name:        et.Name,
			Description: et.Description,
		})
	}

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
	var examVMs []viewmodels.ExamViewModel
	for _, e := range exams {
		examVMs = append(examVMs, viewmodels.ExamViewModel{
			ID:           e.ID,
			PatientName:  e.Patient.FirstName + " " + e.Patient.LastName,
			ExamTypeName: e.ExamType.Name,
			Result:       e.Result,
			CreatedAt:    e.CreatedAt.Format("2006-01-02"),
		})
	}

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

	exams, _ := examRepository.FindAllWithPreloads("Patient", "ExamType")

	var examVMs []viewmodels.ExamViewModel
	for _, e := range exams {
		examVMs = append(examVMs, viewmodels.ExamViewModel{
			ID:           e.ID,
			PatientName:  e.Patient.FirstName + " " + e.Patient.LastName,
			ExamTypeName: e.ExamType.Name,
			Result:       e.Result,
			CreatedAt:    e.CreatedAt.Format("2006-01-02"),
		})
	}

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

	exams, _ := examRepository.FindAll()

	var examVMs []viewmodels.ExamViewModel
	for _, e := range exams {
		examVMs = append(examVMs, viewmodels.ExamViewModel{
			ID:           e.ID,
			PatientName:  e.Patient.FirstName + " " + e.Patient.LastName,
			ExamTypeName: e.ExamType.Description,
			Result:       e.Result,
			CreatedAt:    e.CreatedAt.Format("2006-01-02"),
		})
	}

	return c.Render("exams/exam_list", fiber.Map{
		"Exams": examVMs,
	})
}
