package routes

import (
	"github.com/Amir-Hossein-shamsi/fiber-resapi/src/routes/controllers"
	"github.com/gofiber/fiber/v2"
)

type Connetor struct {
	Handler *controllers.Handller
}

func (c *Connetor) Routes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create", c.Handler.CreateUser)
	api.Get("/getAll", c.Handler.GetUsers)

}
