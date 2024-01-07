package util

import (
	"auto-jj/util/payload"

	"github.com/gofiber/fiber/v2"
)

func Jackson(c *fiber.Ctx, j interface{}, e error) error {
	if e != nil {
		return e
	}
	return c.JSON(payload.ApiResponse{Status: true, Data: j})
}

func JacksonList(c *fiber.Ctx, j interface{}, tot int64, e error) error {
	if e != nil {
		return e
	}
	return c.JSON(payload.ApiListResponse{Status: true, Data: j, TotalCount: tot})
}
