package controllers

import (
	"ChatApp/database"
	"ChatApp/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllAboutme(c *fiber.Ctx) error {

	var aboutmes []models.Aboutme

	database.DB.Find(&aboutmes)

	return c.JSON(aboutmes)
}

func CreateAboutme(c *fiber.Ctx) error {
	var aboutme models.Aboutme

	if err := c.BodyParser(&aboutme); err != nil {
		return err
	}
	database.DB.Create(&aboutme)

	return c.JSON(aboutme)

}

func GetAboutme(c *fiber.Ctx) error {

	type AboutmeWithPosts struct {
		Aboutme models.Aboutme
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	aboutme := models.Aboutme{
		Id: uint(id),
	}

	database.DB.Find(&aboutme)

	database.DB.Preload("User").Table("posts").Select("posts.*, users.*").Joins("JOIN users ON posts.user_id = users.id").Where("users.id = ?", id)

	aboutmeWithPosts := AboutmeWithPosts{
		Aboutme: aboutme,
	}

	return c.JSON(aboutmeWithPosts)
}

func UpdateAboutme(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	aboutme := models.Aboutme{
		Id: uint(id),
	}

	if err := c.BodyParser(&aboutme); err != nil {
		return err
	}
	database.DB.Model(&aboutme).Updates(aboutme)

	return c.JSON(aboutme)
}

func DeleteAboutme(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	aboutme := models.Aboutme{
		Id: uint(id),
	}

	database.DB.Delete(&aboutme)

	return nil
}
