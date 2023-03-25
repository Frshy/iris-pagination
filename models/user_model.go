package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `json:"username"`
}
