package controllers

import (
	"github.com/cocoasterr/net_http/services"
	"github.com/gofiber/fiber/v2"
)

type MutationController struct {
	Mutation services.Mutation
}

func NewMutationController(mutationService services.Mutation) *MutationController {
	return &MutationController{
		Mutation: mutationService,
	}
}
func (con *MutationController) MutationCheck(c *fiber.Ctx) error {
	acc_number := c.Params("account_number")

	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "make sure your input data correct!"})
	}
	ctx := c.Context()
	mutation, err := con.Mutation.IndexMutation(ctx, acc_number)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	for i := range mutation {
		delete(mutation[i], "id")

	}
	resp := map[string]interface{}{
		"status":        "success",
		"data_mutation": mutation,
	}
	return c.JSON(resp)
}
