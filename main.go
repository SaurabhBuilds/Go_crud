package main

import (
	"fmt"
	"log"
	"myapp/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("new fiber application")
	app := fiber.New()
	routers.SetupRouters(app)
	fmt.Println("server is running on port 8000")
	log.Fatal(app.Listen(":8000"))
}
