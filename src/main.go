package main

import (
	"log"

	database "github.com/Amir-Hossein-shamsi/fiber-resapi/src/db"
	"github.com/Amir-Hossein-shamsi/fiber-resapi/src/routes"
	"github.com/Amir-Hossein-shamsi/fiber-resapi/src/routes/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := database.Init()
	h := &controllers.Handller{Handler: database.Repository{DB: db}}
	con := &routes.Connetor{Handler: h}
	con.Routes(app)
	log.Fatal(app.Listen(":4000"))
	log.Println("🤠 Server is Up")
}
