package app

import (
	"GoAPIOnECHO/internal/controller"
	"GoAPIOnECHO/internal/repository"
	"GoAPIOnECHO/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func StartServer() error {
	if err := repository.ConnectToDatabase(repository.DB); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	//I dont know yet how to disable this upon auto creation of table in database
	//Maybe if there is a function for GORM to find table name and if that table name exists
	// dont run this block of code
	//if err := repository.DB.AutoMigrate(&model.Todo{}); err != nil {
	//	log.Fatalf("Failed to auto-migrate models: %v", err)
	//	return err
	//}

	e := echo.New()
	service := &service.TodoService{}

	validator := validator.New()
	ctrl := &controller.Controller{Service: *service, Validator: validator}

	controller.SetupRoutes(e, ctrl)
	if err := e.Start(":8080"); err != nil {
		return echo.NewHTTPError(500, "Something went wrong: "+err.Error())
	}
	return nil
}
