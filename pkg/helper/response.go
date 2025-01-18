package helper

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CheckPassword compares a plaintext password with a hashed password
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func ResponseError(ctx *fiber.Ctx, err error) error {
	var (
		statusCode    int
		customError   string
		originalError string
	)

	statusCode, customError, originalError = TrimMesssage(err)
	fmt.Println(originalError)

	response := Response{
		Message: customError,
		Data:    nil,
	}
	return ctx.Status(statusCode).JSON(response)
}

func ResponseOK(c *fiber.Ctx, msg string, data interface{}) error {
	response := Response{
		Message: msg,
		Data:    data,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func ResponseCreatedOK(c *fiber.Ctx, msg string, data interface{}) error {
	response := Response{
		Message: msg,
		Data:    data,
	}

	return c.Status(http.StatusCreated).JSON(response)
}
