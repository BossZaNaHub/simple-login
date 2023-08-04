package models

import "time"

type ClientMemberProfileResp struct {
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	Email        string    `json:"email"`
	MobileNumber string    `json:"mobile_number"`
	BirthOfDate  time.Time `json:"birthday"`
}
