package helper

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
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
