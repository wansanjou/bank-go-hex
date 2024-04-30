package main

import (
	"fmt"
	"wansanjou/handler"
	"wansanjou/logs"
	"wansanjou/repository"
	"wansanjou/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main() {
	app := fiber.New()
	db := InitializeDB()

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	// customerRepositoryMock := repository.NewCustomerRepositoryMock()
  // _ = customerRepositoryDB
	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	accountRepositoryDB := repository.NewAccountRepositoryDB(db)
	accountService := service.NewAccountService(accountRepositoryDB)
	accountHandler := handler.NewAccountHandler(accountService)

	app.Get("/customers", customerHandler.GetAllCustomer)
	app.Get("/customers/:id", customerHandler.GetCustomer)

	app.Get("/customers/:id/accounts", accountHandler.GetAccounts)
	// app.Post("/customers/:id/accounts" , accountHandler.NewAccount)
	app.Post("/accounts" , accountHandler.NewAccount)

	app.Listen(":8080")
}


func InitializeDB() *gorm.DB {
	const (
		host     = "localhost"  // or the Docker service name if running in another container
		port     = 5432         // default PostgreSQL port
		user     = "myuser"     // as defined in docker-compose.yml
		password = "mypassword" // as defined in docker-compose.yml
		dbname   = "mydatabase" // as defined in docker-compose.yml
	)
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&repository.Customer{})
	db.AutoMigrate(&repository.Account{})

	logs.Info("Starting server at port :5050")
	return db
}