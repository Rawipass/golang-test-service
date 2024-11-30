package routes

import (
	product_http "github.com/Rawipass/golang-test-service/internal/product/http"
	user_http "github.com/Rawipass/golang-test-service/internal/user/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *user_http.UserHandler, productHandler *product_http.ProductHandler) *gin.Engine {
	router := gin.Default()
	//user
	router.GET("/users/list", userHandler.ListUsers)
	router.GET("/user/:id", userHandler.GetUserDetail)
	router.PATCH("/user/deduct/balance/:id", userHandler.DeductBalance)
	router.PATCH("/user/add/balance/:id", userHandler.AddBalance)

	//product
	router.POST("/product", productHandler.CreateProduct)
	router.GET("/product/list", productHandler.ListProducts)
	router.GET("/product/:id", productHandler.GetProductDetail)

	//commission
	router.GET("/commission/:id", productHandler.GetCommissionDetail)
	router.GET("/commission/list", productHandler.ListCommissions)

	//affiliate
	router.POST("/affiliate", productHandler.CreateAffiliate)
	router.GET("/affiliate/list", productHandler.ListAffiliates)
	router.GET("/affiliate/:id", productHandler.GetAffiliateDetail)

	//sale
	router.POST("/sale", productHandler.CreateSale)

	return router
}
