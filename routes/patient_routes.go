package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/handlers"
)

func SetupPatientRoutes(app *fiber.App) {
	app.Get("/", handlers.GetPatients)
	app.Get("/patients", handlers.GetPatients)
	app.Get("/patients/list", handlers.GetPatientsList)
	app.Get("/patients/new", handlers.GetPatientForm)
	app.Post("/patients", handlers.CreatePatient)
	app.Get("/patients/:id", handlers.GetPatient)
	app.Get("/patients/:id/edit", handlers.GetEditPatientForm)
	app.Put("/patients/:id", handlers.UpdatePatient)
	app.Delete("/patients/:id", handlers.DeletePatient)
}
