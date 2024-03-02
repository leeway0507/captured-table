package handlers

import (
	"backend/ent"
	"backend/pkg/store"
	"context"

	"github.com/gofiber/fiber/v2"
)

func GetStores(session *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// param
		ctx := context.Background()

		result, err := store.GetStores(ctx, session)

		if err != nil {
			return HandlerErr(c, err.Error())
		}

		return c.JSON(fiber.Map{"data": result})
	}

}

func GetStore(session *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// param
		ctx := context.Background()
		param := c.Params("id")

		result, err := store.GetStore(ctx, session, param)

		if err != nil {
			return HandlerErr(c, err.Error())
		}

		return c.JSON(fiber.Map{"data": result})
	}

}
