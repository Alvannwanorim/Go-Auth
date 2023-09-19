package models

import "time"

type UserType string

const (
	ADMIN       UserType = "Admin"
	SUPER_ADMIN UserType = "Super_Admin"
	USER        UserType = "User"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"type:varchar(50)" json:"first_name"`
	LastName  string    `gorm:"type:varchar(50)" json:"last_name"`
	Email     string    `gorm:"unique;type:varchar(50)" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"Default:CURRENT_TIMESTAMP" json:"updated_at"`
	UserType  UserType  `gorm:"type:varchar(50);default:'User'" json:"user_type"`
}
