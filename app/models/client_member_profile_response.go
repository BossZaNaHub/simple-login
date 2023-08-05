package models

import "time"

type ClientMemberProfileResp struct {
	UserId       int64     `json:"userId"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	MobileNumber string    `json:"mobile_number"`
	BirthOfDate  time.Time `json:"birthday"`
}
