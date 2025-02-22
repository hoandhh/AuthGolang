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

var jwtKey = []byte("secret_key")
var refreshKey = []byte("refresh_key")

type Credentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Mã hóa mật khẩu
func HashPassword(password string) (string, error) {
	// Tạo một chuỗi băm của password bằng thuật toán Blowfish, DefaultCost=10 là chạy với vòng lặp 2^10 lần
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Kiểm tra mật khẩu
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Kiểm tra email hợp lệ bằng regex
func ValidateEmail(email string) bool {
	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func GenerateTokens(email string) (string, string, error) {
	// Tạo access token
	// Access token hết hạn trong 15 phút
	accessExp := time.Now().Add(15 * time.Minute)

	// Claims chứa thông tin của token, gồm email và thời gian hết hạn của token
	accessClaims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExp.Unix(),
		},
	}

	// Tạo một JWT token mới với payload là accessClaims, ký bằng phương thức SHA256
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	// Ký token với jwtKey, trả về chuỗi JWT dạng header.payload.signature
	accessString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	// ---------------------------------------------------------------------------------
	// Tạo refresh token
	refreshExp := time.Now().Add(7 * 24 * time.Hour)

	// Claims chứa thông tin của token, gồm email và thời gian hết hạn của token
	refreshClaims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExp.Unix(),
		},
	}

	// Tạo một JWT token mới với payload là refreshClaims, ký bằng phương thức SHA256
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// Ký token với jwtKey, trả về chuỗi JWT dạng header.payload.signature
	refreshString, err := refreshToken.SignedString(refreshKey)
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}

func RegisterUser(email, password string) error {
	// Valid email
	if !ValidateEmail(email) {
		return errors.New("email không hợp lệ")
	}

	// Valid password
	if len(password) < 6 {
		return errors.New("mật khẩu phải có ít nhất 6 kí tự")
	}

	// Kiểm tra xem email đã tồn tại chưa
	var existingUser models.User
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return errors.New("email đã được sử dụng")
	}

	// Mã hoá mật khẩu
	hashPassword, err := HashPassword(password)
	if err != nil {
		return errors.New("lỗi khi mã hoá mật khẩu")
	}

	// Tạo user mới và lưu vào DB
	user := models.User{Email: email, Password: hashPassword}
	if err := database.DB.Create(&user).Error; err != nil {
		return errors.New("lỗi không thể tạo tài khoản")
	}

	return nil
}

func LoginUser(email, password string) (string, string, error) {
	var user models.User
	// Truy vấn tìm email trong DB và lấy bản ghi đầu tiên
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", "", errors.New("email không đúng")
	}

	// Kiểm tra password
	if !CheckPasswordHash(password, user.Password) {
		return "", "", errors.New("mật khẩu không đúng")
	}

	// Sinh ra accessToken và refreshToken
	accessToken, refreshToken, err := GenerateTokens(email)
	if err != nil {
		return "", "", errors.New("không thể tạo token")
	}

	return accessToken, refreshToken, nil
}

func RefreshAccessToken(refreshToken string) (string, error) {
	// Giải mã refreshToken
	// Hàm ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error)
	// Hàm callback func(token *jwt.Token) (interface{}, error) trả về refreshKey – khóa bí mật để xác thực token
	// Khi giải mã JWT, ParseWithClaims sẽ ghi dữ liệu vào claims. Nếu ta truyền một giá trị (Claims{}), thì nó sẽ tạo một bản sao, và dữ liệu sẽ không được ghi vào struct gốc
	// bởi jwt.Claims là một interface và interface trong Golang chỉ hoạt động với con trỏ.
	// → Dùng &Claims{} để đảm bảo struct gốc có thể nhận dữ liệu sau khi giải mã.
	token, err := jwt.ParseWithClaims(refreshToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshKey, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("refresh token không hợp lệ")
	}

	// Kiểm tra xem token.Claims có phải là kiểu dữ liệu *Claims(con trỏ đến Claims) hay không.
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("refresh token không hợp lệ")
	}

	// Tạo access token mới
	accessToken, _, err := GenerateTokens(claims.Email)
	if err != nil {
		return "", errors.New("không thể tạo access token mới")
	}

	return accessToken, nil
}
