package routes

import (
	"context"

	"github.com/edanko/moses/internal/entities"
	"github.com/edanko/moses/internal/service/remnant"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getRemnantsAll(service remnant.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		remnants, err := service.GetAll(context.Background())

		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		return c.JSON(&fiber.Map{
			"success":  true,
			"remnants": remnants,
		})
	}
}

func getRemnants(service remnant.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		project := c.Params("project")
		dimension := c.Params("dimension")
		quality := c.Params("quality")

		remnants, err := service.GetNotUsed(context.Background(), project, dimension, quality)

		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"success":  true,
			"remnants": remnants,
		})
	}
}

func addOrUpdateRemnant(service remnant.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(entities.Remnant)
		var result *entities.Remnant
		err := c.BodyParser(requestBody)
		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		if requestBody.ID != primitive.NilObjectID {
			result, err = service.Update(context.Background(), requestBody)
		} else {
			result, err = service.Create(context.Background(), requestBody)
		}
		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"error":   err,
		})
	}
}

func deleteRemnant(service remnant.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dberr := service.Delete(context.Background(), c.Params("id"))
		if dberr != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   dberr.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"message": "deleted successfully",
		})
	}
}

func RemnantRouter(app fiber.Router, service remnant.UseCase) {
	app.Get("/remnants", getRemnantsAll(service))
	app.Get("/remnant/:project/:dimension/:quality", getRemnants(service))
	app.Post("/remnant", addOrUpdateRemnant(service))
	app.Delete("/remnant/:id", deleteRemnant(service))
}
