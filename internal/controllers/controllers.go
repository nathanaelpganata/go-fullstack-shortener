package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nathanaelpganata/go-fullstack-shortener/internal/database"
	"github.com/nathanaelpganata/go-fullstack-shortener/internal/models"
	"github.com/nathanaelpganata/go-fullstack-shortener/internal/utils"
)

func IndexPage(c *fiber.Ctx) error {
	errorMessage := c.Query("errorMessage")
	return c.Render("index", fiber.Map{
		"ErrorMessage": errorMessage,
	})
}

func Redirect(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")

	var link models.Link
	if err := database.DB.Where("short_url = ?", shortUrl).First(&link).Error; err != nil {
		return c.SendString("Failed to redirect")
	}

	if !link.ExpiresAt.IsZero() && time.Now().After(link.ExpiresAt) {
		return c.SendString("Failed to redirect")
	}

	link.Clicks++
	database.DB.Save(&link)

	return c.Redirect(link.OriginalURL)
}

func ShortenPage(c *fiber.Ctx) error {
	DOMAIN_NAME := os.Getenv("DOMAIN_NAME")
	shortUrl := c.Query("shortUrl")

	return c.Render("shorten", fiber.Map{
		"ShortUrl":    shortUrl,
		"DOMAIN_NAME": DOMAIN_NAME,
	})
}

func Shorten(c *fiber.Ctx) error {
	originalUrl := c.FormValue("originalUrl")
	shortUrl := c.FormValue("shortUrl")

	if !utils.IsValidURL(originalUrl) || len(originalUrl) > 255 {
		errorMessage := "Invalid or too long original URL"
		return c.Redirect("/?errorMessage=" + errorMessage)
	}

	if !utils.IsValidShortURL(shortUrl) {
		errorMessage := "Invalid short URL"
		return c.Redirect("/?errorMessage=" + errorMessage)
	}

	var link models.Link
	link.OriginalURL = originalUrl
	link.ShortURL = shortUrl
	link.RequestIP = utils.GetRequestIP(c)
	link.ExpiresAt = time.Now().Add(3 * 24 * time.Hour)

	if err := database.DB.Create(&link).Error; err != nil {
		errorMessage := "Please try again with another short URL"
		return c.Redirect("/?errorMessage=" + errorMessage)
	}

	return c.Redirect("/shorten?shortUrl=" + shortUrl)
}
