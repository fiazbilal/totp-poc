# 🔐 TOTP 2FA POC with Go + Google Authenticator

This is a minimal proof-of-concept (POC) project that demonstrates how to implement Time-based One-Time Passwords (TOTP) for 2FA using **Go**, **Fiber**, and **Google Authenticator**.

Compatible with any TOTP-based authenticator app like:
- Google Authenticator
- Authy
- Microsoft Authenticator

---

## ✨ Features

- 📧 Register a user with email
- 🔑 Generate a TOTP secret and QR code
- 📲 Scan QR code with Google Authenticator
- ✅ Validate 6-digit TOTP code
- 🔒 Fully in-memory (no database) for simplicity

---

## 📦 Tech Stack

- [Go](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [pquerna/otp](https://github.com/pquerna/otp)
- [skip2/go-qrcode](https://github.com/skip2/go-qrcode)

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/your-username/totp-poc.git
cd totp-poc
```

### 2. Initialize Go module

```bash
go mod tidy
```

### 3. Run the app

```bash
go run main.go
```

Server will start at: http://localhost:3000

## Test

### Register

```bash
curl -X POST http://localhost:3000/register \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com"}' | jq -r '.qr_base64' | base64 -d > qr.png
```

### Verify

```bash
curl -X POST http://localhost:3000/verify  
    -H "Content-Type: application/json"
    -d '{"email": "user@example.com", "code": "481432"}'
```