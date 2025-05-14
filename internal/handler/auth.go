package handler

import (
	"errors"
	"net/http"

	"my-project/internal/database"
	"my-project/internal/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthHandler struct {
    db database.Service
}

func NewAuthHandler(db database.Service) *AuthHandler {
    return &AuthHandler{db: db}
}
// HelloAuth handles the GET request for auth root endpoint
func (h *AuthHandler) HelloAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from auth group"})
}

// SignUp handles user registration
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Check if the email is already registered
	var existingUser model.User
	if err := h.db.DB().Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create a new user
	user := &model.User{
		Email:    req.Email,
		Password: string(hashedPassword), // convert []byte to string
	}

	if err := h.db.DB().Create(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User signed up successfully",
		"email":   user.Email,
	})
}





// SignIn handles user authentication
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new user
	user := &model.User{
		Email:    req.Email,
		Password: req.Password, 
	}

	

	c.JSON(http.StatusCreated, gin.H{
		"message": "User signed in successfully",
		"email":   user.Email,
	})
}