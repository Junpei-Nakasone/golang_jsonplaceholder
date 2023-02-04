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

	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
