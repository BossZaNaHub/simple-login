package jwt

type ACToken struct {
	AccessToken        string `json:"access_token"`
	AccessTokenExpire  int64  `json:"access_token_expire"`
	RefreshToken       string `json:"refresh_token"`
	RefreshTokenExpire int64  `json:"refresh_token_expire"`
	Domain             string `json:"domain"`
}
