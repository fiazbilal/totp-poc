package utils

import (
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

func GenerateTOTP(email string) (secret, url string, err error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "TOTP-POC",
		AccountName: email,
	})
	if err != nil {
		return "", "", err
	}
	return key.Secret(), key.URL(), nil
}

func GenerateQRCode(url string) ([]byte, error) {
	return qrcode.Encode(url, qrcode.Medium, 256)
}

func ValidateTOTPCode(secret, code string) bool {
	return totp.Validate(code, secret)
}
