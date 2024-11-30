package http

import (
	"net/http"
	"strconv"

	"github.com/Rawipass/golang-test-service/internal/user/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	useCase *usecase.UserUseCase
}

func NewUserHandler(useCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	users, totalPages, err := h.useCase.ListUsers(limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"page":        page,
		"total_pages": totalPages,
		"users":        users,
	}
	c.JSON(http.StatusOK, gin.H{"Users": response})
}

func (h *UserHandler) GetUserDetail(c *gin.Context) {
	id := c.Param("id")

	user, err := h.useCase.GetUserDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) DeductBalance(c *gin.Context) {
	id := c.Param("id")

	var payload struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := h.useCase.DeductBalance(id, payload.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandler) AddBalance(c *gin.Context) {
	id := c.Param("id")

	var payload struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := h.useCase.AddBalance(id, payload.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
