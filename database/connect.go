package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dataSrc := "hoan:123456@tcp(127.0.0.1:3306)/tempo?charset=utf8mb4&parseTime=True&loc=Local"

	// Mở kết nối đến cơ sở dữ liệu sử dụng GORM và MySQL driver
	db, err := gorm.Open(mysql.Open(dataSrc), &gorm.Config{})

	// Nếu có lỗi xảy ra khi kết nối, ghi log và dừng chương trình
	if err != nil {
		log.Fatal("Không thể kết nối database:", err)
	}

	// Gán kết nối cơ sở dữ liệu cho biến toàn cục DB
	DB = db
	fmt.Println("Kết nối MySQL thành công!")
}
