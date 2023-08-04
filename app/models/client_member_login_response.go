package models

type ClientMemberLoginResp struct {
	Name                string `json:"name"`
	Email               string `json:"email"`
	MobileNumber        string `json:"mobile_number"`
	AccessToken         string `json:"access_token"`
	AccessTokenExpired  int64  `json:"access_token_expired"`
	RefreshToken        string `json:"refresh_token"`
	RefreshTokenExpired int64  `json:"refresh_token_expired"`
}
