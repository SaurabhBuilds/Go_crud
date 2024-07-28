package controllers

import (
	"myapp/models"

	"github.com/gofiber/fiber/v2"
)

func Hellohandler(c *fiber.Ctx) error {
	return c.SendString("Hello Saurabh!")
}

func GetUsers(c *fiber.Ctx) error {
	if len(models.Users) == 0 {
		return c.Status(fiber.StatusOK).JSON(nil)
	}
	return c.JSON(models.Users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("something went wrong")
	}
	models.Users = append(models.Users, user)
	return c.JSON(models.Users)
}

func UpdateUsers(c *fiber.Ctx) error {
	name := c.Params("name")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("something went wrong")
	}
	for index, val := range models.Users {
		if val.Name == name {
			models.Users[index] = user
			return c.JSON(models.Users)
		}
	}
	return c.Status(fiber.StatusBadRequest).SendString("user is present ")
}

func DeleteUser(c *fiber.Ctx) error {
	name := c.Params("name")
	for index, val := range models.Users {
		if val.Name == name {
			models.Users = append(models.Users[:index], models.Users[index+1:]...)
			return c.JSON(models.Users)
		}
	}
	return c.Status(fiber.StatusOK).SendString("user is not present")
}
