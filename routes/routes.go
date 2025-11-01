package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	SetupPatientRoutes(app)
	SetupMedicationRoutes(app)
	SetupPrescriptionRoutes(app)
	SetupExamRoutes(app)
}
