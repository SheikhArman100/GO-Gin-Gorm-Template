package auth

import (
	"net/http"

	"my-project/internal/database"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	db database.Service
}

func NewAuthHandler(db database.Service) *AuthHandler { // Changed to exported function
	return &AuthHandler{
		db: db,
	}
}

// HelloAuth handles the GET request for auth root endpoint
func (h *AuthHandler) HelloAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from auth group"})
}