package utils

import (
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetRequestIP(c *fiber.Ctx) string {
	ip := c.Get("X-Real-IP")
	if ip == "" {
		ip = c.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = c.IP()
	}

	if strings.Contains(ip, ",") {
		ips := strings.Split(ip, ",")
		ip = strings.TrimSpace(ips[0])
	}

	return ip
}

func IsValidURL(url string) bool {
	regex := regexp.MustCompile(`^https?://.+$`)
	return regex.MatchString(url)
}

func IsValidShortURL(shortUrl string) bool {
	// Check the length of shortUrl
	return len(shortUrl) >= 4 && len(shortUrl) <= 50
}
