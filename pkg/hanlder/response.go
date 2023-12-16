package hanlder

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *fiber.Ctx, statusCode int, message string) {
	logrus.Error(message)
	c.JSON(errorResponse{message})
}

type statusResponse struct {
	Status string `json:"status"`
}
