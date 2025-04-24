package main

import (
	"Projects/internal/database"
	"Projects/internal/handlers"
	"Projects/internal/taskService"
	"Projects/internal/userService"
	"Projects/internal/web/tasks"
	"Projects/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()

	taskRepo := taskService.NewTaskRepository(database.DB)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)
	userStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, userStrictHandler)
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
