package controllers

import (
	"context"
	"log"
	"os"

	"github.com/RianIhsan/goCloudinary/database"
	"github.com/RianIhsan/goCloudinary/models"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func User(c *fiber.Ctx) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal akses file .env")
	}

	var user models.User

	_ = c.BodyParser(&user)

	fileHeader, _ := c.FormFile("image")
	file, _ := fileHeader.Open()

	log.Println(fileHeader.Filename)

	ctx := context.Background()

	urlCloudinary := os.Getenv("URLCLOUDINARY")
	cldService, _ := cloudinary.NewFromURL(urlCloudinary)
	response, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	log.Println(response.SecureURL)

	user.Image = response.SecureURL

	database.DB.Create(user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}
