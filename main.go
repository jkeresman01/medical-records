package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/jkeresman01/medical-records/config"
	"github.com/jkeresman01/medical-records/db"
	"github.com/jkeresman01/medical-records/routes"
)

func main() {
	cfg := config.GetFromEnv()

	db.Connect(cfg)
	db.Migrate()

	views := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: views,
	})

	app.Static("/static", "./static")
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
