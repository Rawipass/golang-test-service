package http

import (
	"net/http"

	"github.com/Rawipass/golang-test-service/internal/product/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	useCase *usecase.ProductUseCase
}

func NewProductHandler(useCase *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{useCase: useCase}
}

// Product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var payload struct {
		Name     string  `json:"name"`
		Quantity int     `json:"quantity"`
		Price    float64 `json:"price"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	product, err := h.useCase.CreateProduct(payload.Name, payload.Quantity, payload.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.useCase.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (h *ProductHandler) GetProductDetail(c *gin.Context) {
	id := c.Param("id")

	product, err := h.useCase.GetProductDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

// Commission
func (h *ProductHandler) GetCommissionDetail(c *gin.Context) {
	id := c.Param("id")

	commission, err := h.useCase.GetCommissionDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"commission": commission})
}

func (h *ProductHandler) ListCommissions(c *gin.Context) {
	commissions, err := h.useCase.ListCommissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"commissions": commissions})
}

// Affiliate
func (h *ProductHandler) CreateAffiliate(c *gin.Context) {
	var payload struct {
		Name     string `json:"name"`
		MasterID int    `json:"master_id"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	affiliate, err := h.useCase.CreateAffiliate(payload.Name, payload.MasterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"affiliate": affiliate})
}

func (h *ProductHandler) ListAffiliates(c *gin.Context) {
	affiliates, err := h.useCase.ListAffiliates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"affiliates": affiliates})
}

func (h *ProductHandler) GetAffiliateDetail(c *gin.Context) {
	affiliateID := c.Param("id")

	affiliate, err := h.useCase.GetAffiliateDetail(affiliateID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"affiliates": affiliate})
}

//sale

func (h *ProductHandler) CreateSale(c *gin.Context) {
	var request struct {
		AffiliateID  uuid.UUID `json:"affiliate_id"`
		ProductID    uuid.UUID `json:"product_id"`
		ProductPrice float64   `json:"product_price"`
	}

	// Bind the JSON request body to the struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the usecase to create the sale and calculate commission
	commissions, err := h.useCase.CreateSale(request.AffiliateID, request.ProductID, request.ProductPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"commissions": commissions})
}
