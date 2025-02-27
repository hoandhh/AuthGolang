package models

type User struct {
	// ID không âm, tự động tăng
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
