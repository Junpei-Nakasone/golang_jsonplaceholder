package main

import (
	"fmt"
	"golang_jsonplaceholder_without_framework/config"
	"golang_jsonplaceholder_without_framework/controller"
	"golang_jsonplaceholder_without_framework/helper"
	"golang_jsonplaceholder_without_framework/repository"
	"golang_jsonplaceholder_without_framework/router"
	"golang_jsonplaceholder_without_framework/service"
	"net/http"
)

func main() {
	fmt.Printf("Start server")

	db := config.DatabaseConnection()

	bookRepository := repository.NewBookRepository(db)

	bookService := service.NewBookServiceImpl(bookRepository)

	bookController := controller.NewBookController(bookService)

	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserServiceImpl(userRepository)

	userController := controller.NewUserController(userService)

	routes := router.NewRouter(bookController, userController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
