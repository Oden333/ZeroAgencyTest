package hanlder

import (
	"ZAtest/models"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (h *Handler) ListNewsHandler(c *fiber.Ctx) error {
	// Обработка запроса на получение списка новостей
	// Вызов функции для запроса данных из базы данных
	// Формирование и возврат ответа в формате JSON
	return c.JSON(fiber.Map{
		"Success": true,
		"News":    []fiber.Map{
			// Данные новостей
		},
	})
}

func (h *Handler) EditNewsHandler(c *fiber.Ctx) error {
	var newsData models.News

	// Извлечение данных из запроса.
	if err := c.BodyParser(&newsData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID not fou"})
	}

	// Извлечение идентификатора из параметра URL.
	id := c.Params("Id")

	// Преобразование идентификатора в int64
	newsID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid news ID"})
	}

	logrus.Println(newsData)
	// Вызов метода сервиса для редактирования новости.
	if err := h.services.EditNewsById(newsID, newsData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to edit news"})
	}

	return c.JSON(fiber.Map{
		"Success": true,
		"Id":      id,
	})
}

func (h *Handler) GetAllNewsHandler(c *fiber.Ctx) error {
	// Вызов метода сервиса для получения новостей.
	list, err := h.services.GetAllNews()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return c.JSON(fiber.Map{
			"Success": false,
			"News":    err,
		})
	}
	return c.JSON(fiber.Map{
		"Success": true,
		"News":    list,
	})
}
