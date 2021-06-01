package controllers

import (
	"crypto/rand"
	"go-trendy-wash-backend/db"
	"math"
	"time"
)

var (
	loc, loc_err = time.LoadLocation("Asia/Bangkok")
	timeLog      = time.Now().In(loc)
	timeInc      = time.Hour * 7
	conn         = db.ConnectDB()
)

func weekday(day int) string {
	switch day {
	case 0:
		return "sunday"
	case 1:
		return "monday"
	case 2:
		return "tuesday"
	case 3:
		return "wednesday"
	case 4:
		return "thursday"
	case 5:
		return "friday"
	case 6:
		return "saturday"
	}
	return "sunday"
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

//GenerateRandomString : ..
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	// const letters = "123456789"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

func paymentEnum(input string) string {
	switch input {
	case "qr_test":
		return "QR Code"
	case "tw_pay":
		return "QR Code"
	case "linepay":
		return "Line Pay"
	case "truemoney":
		return "Truemoney Wallet"
	case "omise":
		return "Credit Card"
	case "alipay":
		return "Alipay"
	case "scb":
		return "Thai QR-Code"
	case "admin":
		return "Admin"
	default:
		return input
	}
}

func machineStatusEnum(input string) string {
	switch input {
	case "0":
		return "Running"
	default:
		return "Loss Connection"
	}
}

func isNan(input float64) float64 {
	if math.IsNaN(input) || input == math.Inf(-1) || input == math.Inf(1) {
		return 0.0
	}
	return input
}

func checkIntString(input string) string {
	if input == "" {
		return "0"
	}
	return input
}
