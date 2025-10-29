package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/repository/factory"
	"github.com/jkeresman01/medical-records/viewmodels"
)

func GetExams(c *fiber.Ctx) error {
	examRepo := repositoryfactory.GetInstance[models.Exam]()
	patientRepo := repositoryfactory.GetInstance[models.Patient]()
	examTypeRepo := repositoryfactory.GetInstance[models.ExamType]()

	exams, err := examRepo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching exams")
	}

	var examVMs []viewmodels.ExamViewModel
	for _, e := range exams {
		patient, _ := patientRepo.FindByID(e.PatientID)
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

	examTypes, _ := examTypeRepo.FindAll()
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
	patientRepo := repositoryfactory.GetInstance[models.Patient]()
	examTypeRepo := repositoryfactory.GetInstance[models.ExamType]()

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

	examTypes, _ := examTypeRepo.FindAll()
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
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	examRepo := repositoryfactory.GetInstance[models.Exam]()
	patientRepo := repositoryfactory.GetInstance[models.Patient]()
	examTypeRepo := repositoryfactory.GetInstance[models.ExamType]()

	exam, err := examRepo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Exam not found")
	}

	patient, _ := patientRepo.FindByID(exam.PatientID)
	examType, _ := examTypeRepo.FindByID(exam.ExamTypeID)

	examVM := viewmodels.ExamViewModel{
		ID:           exam.ID,
		PatientID:    exam.PatientID,
		ExamTypeID:   exam.ExamTypeID,
		PatientName:  patient.FirstName + " " + patient.LastName,
		ExamTypeName: examType.Name,
		Result:       exam.Result,
		CreatedAt:    exam.CreatedAt.Format("2006-01-02"),
	}

	patients, _ := patientRepo.FindAll()
	examTypes, _ := examTypeRepo.FindAll()

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
		"Exam":      examVM,
		"Patients":  patientVMs,
		"ExamTypes": examTypeVMs,
	})
}

func CreateExam(c *fiber.Ctx) error {
	examRepo := repositoryfactory.GetInstance[models.Exam]()
	patientRepo := repositoryfactory.GetInstance[models.Patient]()
	examTypeRepo := repositoryfactory.GetInstance[models.ExamType]()

	patientID, _ := strconv.ParseUint(c.FormValue("patient_id"), 10, 32)
	examTypeID, _ := strconv.ParseUint(c.FormValue("exam_type_id"), 10, 32)

	exam := &models.Exam{
		PatientID:  uint(patientID),
		ExamTypeID: uint(examTypeID),
		Result:     c.FormValue("result"),
	}

	if err := examRepo.Create(exam); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating exam")
	}

	exams, _ := examRepo.FindAll()

	var examVMs []viewmodels.ExamViewModel
	for _, e := range exams {
		patient, _ := patientRepo.FindByID(e.PatientID)
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

	return c.Render("exams/exam_list", fiber.Map{
		"Exams": examVMs,
	})
}

func UpdateExam(c *fiber.Ctx) error {
	examRepo := repositoryfactory.GetInstance[models.Exam]()
	patientRepo := repositoryfactory.GetInstance[models.Patient]()
	examTypeRepo := repositoryfactory.GetInstance[models.ExamType]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	exam, err := examRepo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Exam not found")
	}

	patientID, _ := strconv.ParseUint(c.FormValue("patient_id"), 10, 32)
	examTypeID, _ := strconv.ParseUint(c.FormValue("exam_type_id"), 10, 32)

	exam.PatientID = uint(patientID)
	exam.ExamTypeID = uint(examTypeID)
	exam.Result = c.FormValue("result")

	if err := examRepo.Update(exam); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating exam")
	}

	exams, _ := examRepo.FindAll()

	var examVMs []viewmodels.ExamViewModel
	for _, e := range exams {
		patient, _ := patientRepo.FindByID(e.PatientID)
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

	return c.Render("exams/exam_list", fiber.Map{
		"Exams": examVMs,
	})
}

func DeleteExam(c *fiber.Ctx) error {
	examRepo := repositoryfactory.GetInstance[models.Exam]()
	patientRepo := repositoryfactory.GetInstance[models.Patient]()
	examTypeRepo := repositoryfactory.GetInstance[models.ExamType]()

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	if err := examRepo.DeleteByID(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting exam")
	}

	exams, _ := examRepo.FindAll()

	var examVMs []viewmodels.ExamViewModel
	for _, e := range exams {
		patient, _ := patientRepo.FindByID(e.PatientID)
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

	return c.Render("exams/exam_list", fiber.Map{
		"Exams": examVMs,
	})
}
