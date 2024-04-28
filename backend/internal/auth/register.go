package auth

import (
	"net/http"

	"intelli-learn/backend/internal/database"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(ctx *gin.Context) {
	var request RegisterRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if request.Email == "" || request.Username == "" || request.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Fields cannot be left empty"})
		return
	}

	_, err := database.SelectUserByEmail(request.Email)
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	_, err = database.SelectUserByUsername(request.Username)
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	err = database.InsertUser(request.Email, request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
