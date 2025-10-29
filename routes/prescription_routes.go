package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/medical-records/handlers"
)

func SetupPrescriptionRoutes(app *fiber.App) {
	app.Get("/prescriptions", handlers.GetPrescriptions)
	app.Get("/prescriptions/new", handlers.GetPrescriptionForm)
	app.Post("/prescriptions", handlers.CreatePrescription)
	app.Get("/prescriptions/:id/edit", handlers.GetEditPrescriptionForm)
	app.Put("/prescriptions/:id", handlers.UpdatePrescription)
	app.Delete("/prescriptions/:id", handlers.DeletePrescription)
}
