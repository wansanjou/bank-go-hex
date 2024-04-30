package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"wansanjou/errs"
	"wansanjou/service"

	"github.com/gofiber/fiber/v2"
)

type accountHandler struct {
	accountservice service.AccountService
}

func NewAccountHandler(accountservice service.AccountService) accountHandler  {
	return accountHandler{accountservice: accountservice}
}

func (ah *accountHandler) NewAccount(c *fiber.Ctx) error {
	customerID, err := strconv.Atoi(c.Params("customerID"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var account service.NewAccountRequest
	if err := c.BodyParser(&account); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	response, err := ah.accountservice.NewAccount(customerID, account)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}


func (ah accountHandler) GetAccounts(c *fiber.Ctx) error  {
	customerID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	account, err := ah.accountservice.GetAccounts(customerID)
	if err != nil {	
		if err == errs.NewNotFoundError("Account not found") {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Account not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(account)
}
