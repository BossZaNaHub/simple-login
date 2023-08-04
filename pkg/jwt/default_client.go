package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/kz-login/env"
	"github.com/kz-login/pkg/errors"
	"log"
	"time"
)

type defaultClient struct {
	cfg *env.Environment
	tz  *time.Location
}

func NewClient(cfg *env.Environment) Client {
	tz, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatal("error load location failed", err)
	}
	return &defaultClient{cfg: cfg, tz: tz}
}

func (c *defaultClient) CreateToken(data ClaimTokenData) (*ACToken, errors.Error) {
	accessTokenTime := time.Now().In(c.tz).Add(time.Minute * time.Duration(c.cfg.JWT.JwtExpirationTime)).Unix()
	refreshTokenTime := time.Now().In(c.tz).Add(time.Minute * time.Duration(c.cfg.JWT.JwtRefreshExpirationTime)).Unix()

	acToken, err := c.generateTokenClaim("", accessTokenTime)
	if err != nil {
		return nil, err
	}
	rfToken, err := c.generateTokenClaim("", refreshTokenTime)
	if err != nil {
		return nil, err
	}

	return &ACToken{
		AccessToken:        acToken,
		AccessTokenExpire:  accessTokenTime,
		RefreshToken:       rfToken,
		RefreshTokenExpire: refreshTokenTime,
	}, nil
}

func (c *defaultClient) generateTokenClaim(uid string, exp int64) (string, errors.Error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = c.cfg.JWT.JwtIssuer
	claims["sub"] = uid
	claims["exp"] = exp
	claims["iat"] = time.Now().In(c.tz).Unix()

	secretKey := []byte(c.cfg.JWT.JwtSecret)
	t, err := token.SignedString(secretKey)
	if err != nil {
		return "", errors.ErrClientTokenInvalid
	}

	return t, nil
}
