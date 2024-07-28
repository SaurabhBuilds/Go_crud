package routers

import (
	"myapp/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {
	app.Get("/", controllers.Hellohandler)
	app.Get("/users", controllers.GetUsers)
	app.Post("/users", controllers.CreateUser)
	app.Put("/user/:name", controllers.UpdateUsers)
	app.Delete("/user/:name", controllers.DeleteUser)

}
