package migrations

import "time"

type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FullName  string    `gorm:"column:full_name;size:255;not null" json:"full_name" binding:"required"`
	Email     string    `gorm:"column:email;size:255;not null" json:"email" binding:"required,email"`
	Phone     string    `gorm:"column:phone;size:255;not null" json:"phone" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}
