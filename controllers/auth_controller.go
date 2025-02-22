package controllers

import (
	"net/http"

	"Auth/services"

	"Auth/dtos"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// Đăng ký tài khoản
func (ctrl AuthController) Register(c *gin.Context) {
	var userDTO dtos.UserDTO
	// Đọc dữ liệu JSON từ request body vào biến userDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		// gin.H là một map[string]any, giúp tạo một JSON response một cách dễ dàng và ngắn gọn
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	err := services.RegisterUser(userDTO.Email, userDTO.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tạo tài khoản thành công"})
}

// Đăng nhập, trả về access token và refresh token
func (ctrl AuthController) Login(c *gin.Context) {
	var userDTO dtos.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	accessToken, refreshToken, err := services.LoginUser(userDTO.Email, userDTO.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Làm mới Access Token
func (ctrl AuthController) RefreshToken(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Thiếu refresh token"})
		return
	}

	newAccessToken, err := services.RefreshAccessToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}
