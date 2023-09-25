package main

import (
	"log"

	"github.com/cocoasterr/net_http/app/controllers"
	PGConfig "github.com/cocoasterr/net_http/infra/db/postgres"
	PGRepository "github.com/cocoasterr/net_http/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/net_http/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db_dsn, _ := PGConfig.CreateDatabase("go_bank")
	db, err := PGConfig.ConnectPG(db_dsn)
	if err != nil {
		log.Fatal("Failed to Connect DB")
	}
	if err := PGConfig.CreateTable(db); err != nil {
		log.Fatal("Failed to Create Table!")
	}
	app := fiber.New()
	MutationRepo := PGRepository.NewMutationRepository(db)
	MutationService := services.NewMutationService(MutationRepo.Repository)
	MutationController := controllers.NewMutationController(*MutationService)

	BankRepo := PGRepository.NewUserRepository(db)
	BankService := services.NewAuthService(BankRepo.Repository)
	BankController := controllers.NewAuthController(*BankService, *MutationService)

	app.Post("/daftar", BankController.Register)
	app.Get("/saldo/:account_number", BankController.CheckBalance)
	app.Post("/tabung", BankController.Tabung)
	app.Post("/tarik", BankController.Tarik)

	app.Get("/mutasi/:account_number", MutationController.MutationCheck)

	log.Fatal(app.Listen(":3000"))
}
