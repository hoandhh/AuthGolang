package models

type Employee struct {
	ID      uint   `gorm:"primaryKey"`
	Email   string `gorm:"unique;not null"`
	Name    string `gorm:"not null"`
	Age     int    `gorm:"not null"`
	Address string `gorm:"not null"`
}
