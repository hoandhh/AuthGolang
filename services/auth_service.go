package services

import (
	"errors"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"Auth/database"
	"Auth/models"
)

// jwtKey là khóa bí mật dùng để ký JWT
var jwtKey = []byte("secret_key")

// Credentials là cấu trúc chứa thông tin đăng nhập của người dùng
type Credentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Claims là cấu trúc chứa thông tin của JWT
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Mã hóa mật khẩu
func HashPassword(password string) (string, error) {
	// Sử dụng bcrypt để mã hóa mật khẩu với chi phí mặc định
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Kiểm tra mật khẩu
func CheckPasswordHash(password, hash string) bool {
	// So sánh mật khẩu đã mã hóa với mật khẩu người dùng nhập vào
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Kiểm tra email hợp lệ bằng regex
func ValidateEmail(email string) bool {
	// Biểu thức chính quy để kiểm tra định dạng email
	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// Xử lý đăng ký
func RegisterUser(email, password string) error {
	// Kiểm tra email hợp lệ
	if !ValidateEmail(email) {
		return errors.New("email không hợp lệ")
	}

	// Kiểm tra độ dài mật khẩu
	if len(password) < 6 {
		return errors.New("mật khẩu phải có ít nhất 6 ký tự")
	}

	// Mã hóa mật khẩu
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return errors.New("lỗi mã hóa mật khẩu")
	}

	// Tạo người dùng mới và lưu vào cơ sở dữ liệu
	user := models.User{Email: email, Password: hashedPassword}
	if err := database.DB.Create(&user).Error; err != nil {
		return errors.New("lỗi tạo tài khoản")
	}

	return nil
}

// Xử lý đăng nhập
func LoginUser(email, password string) (string, error) {
	var user models.User
	// Tìm người dùng trong cơ sở dữ liệu theo email
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("sai email hoặc mật khẩu")
	}

	// Kiểm tra mật khẩu
	if !CheckPasswordHash(password, user.Password) {
		return "", errors.New("sai email hoặc mật khẩu")
	}

	// Token hết hạn sau 15 phút
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Tạo JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("không thể tạo JWT")
	}

	return tokenString, nil
}
