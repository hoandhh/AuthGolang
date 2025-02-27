package middleware

import (
	"Auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy token từ header "Authorization" của request
		token := c.GetHeader("Authorization")

		// Gọi hàm ValidateAccessToken để xác thực token và lấy userID
		userID, err := services.ValidateAccessToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			// Dừng ngay việc xử lý request, không đi tiếp đến các handler khác
			// Mỗi handler chịu trách nhiệm xử lý một request HTTP đến một endpoint (đường dẫn) cụ thể
			// Handler nhận request từ client, xử lý logic nghiệp vụ, và trả về response
			c.Abort()
			return
		}

		// Lưu userID vào context để các handler tiếp theo có thể sử dụng
		c.Set("userID", userID)

		// Cho phép request tiếp tục đến handler tiếp theo
		c.Next()
	}
}
