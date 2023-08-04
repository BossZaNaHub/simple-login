package daos

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	MobileNumber      string    `gorm:"column:mobile_number;type:varchar(10);uniqueIndex"`
	Email             string    `gorm:"column:email;type:varchar(255)"`
	Firstname         string    `gorm:"column:first_name;type:varchar(64)"`
	Lastname          string    `gorm:"column:last_name;type:varchar(64)"`
	Birthday          time.Time `gorm:"column:birthday;type:date;not null"`
	IsActive          bool      `gorm:"column:is_active;type:bool;default(false)"`
	PasswordEncrypted string    `gorm:"column:password_encrypted;type:varchar(60);not null"`
}
