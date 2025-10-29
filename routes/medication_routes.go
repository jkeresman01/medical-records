package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/handlers"
)

func SetupMedicationRoutes(app *fiber.App) {
	app.Get("/medications", handlers.GetMedications)
	app.Get("/medications/new", handlers.GetMedicationForm)
	app.Post("/medications", handlers.CreateMedication)
	app.Get("/medications/:id/edit", handlers.GetEditMedicationForm)
	app.Put("/medications/:id", handlers.UpdateMedication)
	app.Delete("/medications/:id", handlers.DeleteMedication)
}
