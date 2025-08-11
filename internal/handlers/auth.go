package handlers

import (
	"net/http"
	"user-auth-service/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.InputRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Validate input
	if input.Password == "" || input.Email == "" {
		c.JSON(400, gin.H{"error": "Password, and email are required"})
		return
	}

	// Call the service to sign up the user
	if token, err := h.service.Authorization.SignUp(input); err != nil {
		c.JSON(500, gin.H{"error": "Failed to sign up"})
		return
	}


	// Return the token in the response
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var input models.InputRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if _, err := h.service.Authorization.Login(input); err != nil {
		c.JSON(500, gin.H{"error": "Failed to login"})
		return
	}

	// Validate input
	if input.Password == "" || input.Email == "" {
		c.JSON(400, gin.H{"error": "Password and email are required"})
		return
	}

	user := models.InputRequest(input)

	//login and generate token
	token, err := h.service.Login(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to login"})
		return
	}

	// Return the token in the response
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
