package hanlder

import (
	"ZAtest/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	app.Post("/edit/:Id", h.EditNewsHandler)
	app.Get("/list", h.GetAllNewsHandler)
}
