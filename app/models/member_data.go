package models

import "time"

type MemberData struct {
	Id                int64     `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	MobileNumber      string    `json:"mobile_number"`
	BirthOfDate       time.Time `json:"birthday"`
	PasswordEncrypted string    `json:"password_encrypted"`
	IsActive          bool      `json:"is_active"`
}
