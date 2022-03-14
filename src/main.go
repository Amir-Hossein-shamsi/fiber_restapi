package main

import (
	"log"

	_ "github.com/Amir-Hossein-shamsi/fiber-resapi/src/db"
	r "github.com/Amir-Hossein-shamsi/fiber-resapi/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	r.Routes(app)
	log.Fatal(app.Listen(":4000"))
	log.Println("🤠 Server is Up")
}
