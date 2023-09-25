package controllers

import (
	"fmt"
	"reflect"

	"github.com/cocoasterr/net_http/services"
	"github.com/gofiber/fiber/v2"
)

type BankController struct {
	AuthService services.BankService
	Mutation    services.Mutation
}

func NewAuthController(authService services.BankService, mutation services.Mutation) *BankController {
	return &BankController{
		AuthService: authService,
		Mutation:    mutation,
	}
}
func (con *BankController) Register(c *fiber.Ctx) error {

	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "make sure your input data correct!"})
	}

	ctx := c.Context()

	conditioQuery := fmt.Sprintf("where phone_number = '%s' or nik = '%s'", payload["phone_number"], payload["nik"])
	existUser, _ := con.AuthService.FindByCustomConditionService(ctx, conditioQuery)
	if existUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "NIK or Phone Number already registered!"})
	}

	if err := con.AuthService.CreateService(ctx, payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	data, _ := con.AuthService.FindByService(ctx, "nik", payload["nik"])

	return c.JSON(fiber.Map{"message": "success!", "Account Number": data["account_number"]})
}

func (con *BankController) CheckBalance(c *fiber.Ctx) error {
	acc_number := c.Params("account_number")

	ctx := c.Context()
	user_data, err := con.AuthService.FindByService(ctx, "account_number", acc_number)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	if user_data == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Account Numbers not found!"})
	}
	resp := map[string]interface{}{
		"status": "success",
		"saldo":  user_data["balance"],
	}
	return c.JSON(resp)
}

func (con *BankController) Tabung(c *fiber.Ctx) error {
	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "make sure your input data correct!"})
	}
	for k, v := range payload {
		if reflect.TypeOf(v) == reflect.TypeOf(float64(0)) {
			payload[k] = int(v.(float64))
		}
	}
	ctx := c.Context()
	userData, err := con.AuthService.FindByService(ctx, "account_number", payload["account_number"])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	if userData == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Account Numbers not found!"})
	}
	transaction := con.AuthService.Transaction(userData, payload["nominal"].(int))
	if err := con.AuthService.UpdateService(ctx, transaction, userData["id"].(string)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	// Create record Mutation
	mutation := map[string]interface{}{"account_number": transaction["account_number"], "transaction": payload["nominal"]}
	if err := con.Mutation.CreateService(ctx, mutation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	resp := map[string]interface{}{
		"status": "success",
		"saldo":  transaction["balance"],
	}
	return c.JSON(resp)
}

func (con *BankController) Tarik(c *fiber.Ctx) error {
	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "make sure your input data correct!"})
	}
	for k, v := range payload {
		if reflect.TypeOf(v) == reflect.TypeOf(float64(0)) {
			payload[k] = int(v.(float64))
		}
	}
	ctx := c.Context()
	userData, err := con.AuthService.FindByService(ctx, "account_number", payload["account_number"])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	if userData == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Account Numbers not found!"})
	}
	nominal := -payload["nominal"].(int)
	transaction := con.AuthService.Transaction(userData, nominal)
	if transaction == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Your balance is not enough!"})
	}
	if err := con.AuthService.UpdateService(ctx, transaction, userData["id"].(string)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	// Create record Mutation
	mutation := map[string]interface{}{"account_number": transaction["account_number"], "transaction": nominal}
	if err := con.Mutation.CreateService(ctx, mutation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	resp := map[string]interface{}{
		"status": "success",
		"saldo":  transaction["balance"],
	}
	return c.JSON(resp)
}
