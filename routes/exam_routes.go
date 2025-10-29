package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/handlers"
)

func SetupExamRoutes(app *fiber.App) {
	app.Get("/exams", handlers.GetExams)
	app.Get("/exams/new", handlers.GetExamForm)
	app.Post("/exams", handlers.CreateExam)
	app.Get("/exams/:id/edit", handlers.GetEditExamForm)
	app.Put("/exams/:id", handlers.UpdateExam)
	app.Delete("/exams/:id", handlers.DeleteExam)
}
