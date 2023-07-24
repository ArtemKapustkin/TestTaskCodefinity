package handler

import (
	"github.com/gofiber/fiber/v2"
)

type FizzBuzzService interface {
	FizzBuzz(n int) *[]string
}

type FizzBuzzHandler struct {
	fizzBuzzService FizzBuzzService
}

func NewFizzBuzzHandler(fizzBuzzService FizzBuzzService) *FizzBuzzHandler {
	return &FizzBuzzHandler{
		fizzBuzzService: fizzBuzzService,
	}
}

type RequestBody struct {
	N int `json:"n"`
}

func (h *FizzBuzzHandler) GetStringArr(c *fiber.Ctx) error {
	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	result := h.fizzBuzzService.FizzBuzz(requestBody.N)

	return c.JSON(result)
}
