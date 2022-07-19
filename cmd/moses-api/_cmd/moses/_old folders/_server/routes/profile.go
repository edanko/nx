package routes

import (
	"context"

	"github.com/edanko/moses/internal/entities"

	"github.com/edanko/moses/internal/service/profile"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getProfilesAll(service profile.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		profiles, err := service.GetAll(context.Background())

		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		return c.JSON(&fiber.Map{
			"success":  true,
			"profiles": profiles,
		})
	}
}

func getProfiles(service profile.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		project := c.Params("project")
		dimension := c.Params("dimension")
		quality := c.Params("quality")

		profiles, err := service.Get(context.Background(), project, dimension, quality)

		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"success":  true,
			"profiles": profiles,
		})
	}
}

func addOrUpdateProfile(service profile.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(entities.Profile)
		var result *entities.Profile
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

func deleteProfile(service profile.UseCase) fiber.Handler {
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

func ProfileRouter(app fiber.Router, service profile.UseCase) {
	app.Get("/profiles", getProfilesAll(service))
	app.Get("/profile/:project/:dimension/:quality", getProfiles(service))
	app.Post("/profile", addOrUpdateProfile(service))
	app.Delete("/profile/:id", deleteProfile(service))
}
