package handler

import (
	"net/http"
	"strconv"
	"wansanjou/errs"
	"wansanjou/service"

	"github.com/gofiber/fiber/v2"
)

type customerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) customerHandler  {
	return customerHandler{customerService: customerService}
}


func (ch customerHandler) GetAllCustomer(c *fiber.Ctx) error  {
	customers , err := ch.customerService.GetAllCustomer() 
	if err != nil {
		return err
	}

	return c.JSON(customers)
}

func (ch customerHandler) GetCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	customer, err := ch.customerService.GetCustomer(id)
	if err != nil {	
		return errs.NewNotFoundError("Customer not found")
	}

	return c.JSON(customer)
}