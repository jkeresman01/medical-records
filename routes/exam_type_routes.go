package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/handlers"
)

func SetupExamTypeRoutes(app *fiber.App) {
	app.Get("/exam-types", handlers.GetExamTypes)
	app.Get("/exam-types/new", handlers.GetExamTypeForm)
	app.Post("/exam-types", handlers.CreateExamType)
	app.Get("/exam-types/:id/edit", handlers.GetEditExamTypeForm)
	app.Put("/exam-types/:id", handlers.UpdateExamType)
	app.Delete("/exam-types/:id", handlers.DeleteExamType)
}
