package jwt

type ClaimTokenData struct {
	UID          int64  `json:"uid"`
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
}
