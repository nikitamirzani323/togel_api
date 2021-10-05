package controller

import (
	"bitbucket.org/isbtotogroup/api_go/model"
	"github.com/gofiber/fiber/v2"
)

type clienrequest struct {
	Tipe string `json:"tipe"`
	Nama string `json:"nama"`
}

func FetchAll_bukumimpi(c *fiber.Ctx) error {
	client := new(clienrequest)
	if err := c.BodyParser(client); err != nil {
		return err
	}

	result, err := model.FetchAll_Mbukumimpi(client.Tipe, client.Nama)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
