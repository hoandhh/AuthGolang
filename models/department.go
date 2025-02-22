package models

type Department struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}
