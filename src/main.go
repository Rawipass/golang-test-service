package main

import (
	"log"

	"github.com/Rawipass/golang-test-service/config"
	product_http "github.com/Rawipass/golang-test-service/internal/product/http"
	product_repository "github.com/Rawipass/golang-test-service/internal/product/repository"
	product_usecase "github.com/Rawipass/golang-test-service/internal/product/usecase"
	user_http "github.com/Rawipass/golang-test-service/internal/user/http"
	user_repository "github.com/Rawipass/golang-test-service/internal/user/repository"
	user_usecase "github.com/Rawipass/golang-test-service/internal/user/usecase"
	"github.com/Rawipass/golang-test-service/routes"
)

func main() {
	// Init Config
	config.InitConfig()

	// Init Database
	config.ConnectDatabase()

	userRepo := user_repository.NewUserRepository(config.DB)
	userUseCase := user_usecase.NewUserUseCase(userRepo)
	userHandler := user_http.NewUserHandler(userUseCase)

	productRepo := product_repository.NewProductRepository(config.DB)
	productUseCase := product_usecase.NewProductUseCase(productRepo)
	productHandler := product_http.NewProductHandler(productUseCase)

	// Setup Route
	r := routes.SetupRouter(userHandler, productHandler)
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Could not run server: %v\n", err)
	}

}
