package handler

import (
	"encoding/base64"
	"totp-poc/utils"

	"github.com/gofiber/fiber/v2"
)

var userSecrets = make(map[string]string) // email -> secret

type RegisterRequest struct {
	Email string `json:"email"`
}

type VerifyRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func RegisterHandler(c *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	secret, url, err := utils.GenerateTOTP(req.Email)
	if err != nil {
		return c.Status(500).SendString("Error generating TOTP")
	}

	qr, err := utils.GenerateQRCode(url)
	if err != nil {
		return c.Status(500).SendString("Error generating QR code")
	}

	userSecrets[req.Email] = secret

	return c.JSON(fiber.Map{
		"secret":    secret,
		"qr_base64": base64.StdEncoding.EncodeToString(qr),
	})
}

func VerifyHandler(c *fiber.Ctx) error {
	req := new(VerifyRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	secret, ok := userSecrets[req.Email]
	if !ok {
		return c.Status(404).SendString("User not found")
	}

	if utils.ValidateTOTPCode(secret, req.Code) {
		return c.SendString("✅ Code is valid!")
	}
	return c.Status(401).SendString("❌ Invalid code.")
}
