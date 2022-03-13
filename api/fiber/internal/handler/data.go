package handler

import (
	"encoding/csv"

	"github.com/gofiber/fiber/v2"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/usecases/data"
)

const (
	NUM_WORKER       = 10
	WRITE_BATCH_SIZE = 1000
)

func NewDataHandler(app fiber.Router, service data.DataUseCase) {
	app.Post("/upload", uploadCsv(service))
}

func uploadCsv(service data.DataUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}
		fileBuffer, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		r := csv.NewReader(fileBuffer)
		lines, err := r.ReadAll()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		inserted, err := service.CreateDatasConcurrent(lines, WRITE_BATCH_SIZE, NUM_WORKER)

		return c.JSON(&fiber.Map{
			"status":   "success",
			"error":    err,
			"inserted": inserted,
		})
	}
}
