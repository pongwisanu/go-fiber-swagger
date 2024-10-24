package main

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

type Random struct {
	Value int `json:"value"`
}

type Error struct {
	Error string `json:"error"`
}

// Random godoc
// @Summary Random Number
// @Description Random number by Max
// @Tags number
// @Accept json
// @Produce json
// @Param max path int true "Max"
// @Success 200 {object} Random
// @Failure 406 {object} Error
// @Router /random/{max} [get]
func RandomNumber(c *fiber.Ctx) error {

	max, err := c.ParamsInt("max")
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(Error{
			Error: err.Error(),
		})
	}

	rand := 1 + rand.Intn(max-0)

	return c.JSON(Random{
		Value: rand,
	})
}
