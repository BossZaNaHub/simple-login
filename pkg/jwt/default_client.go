package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kz-login/env"
	"github.com/kz-login/pkg/errors"
	"log"
	"strconv"
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

	acToken, err := c.generateTokenClaim(data.UID, accessTokenTime)
	if err != nil {
		return nil, err
	}
	rfToken, err := c.generateTokenClaim(data.UID, refreshTokenTime)
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

func (c *defaultClient) RefreshToken(rfToken string) (*ACToken, errors.Error) {
	token, err := jwt.Parse(rfToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.cfg.JWT.JwtSecret), nil
	})

	log.Print(token, err)
	if err != nil || !token.Valid {
		return nil, errors.NewError(errors.ErrCodeClientTokenInvalid, "invalid or expired refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.NewError(errors.ErrCodeClientTokenInvalid, "failed to parse refresh token")
	}

	u, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.NewError(errors.ErrCodeClientTokenInvalid, "failed to extract uid")
	}

	uid, err := strconv.Atoi(u)
	if err != nil {
		return nil, errors.NewError(errors.ErrCodeClientTokenInvalid, "failed to extract uid")
	}

	accessTokenTime := time.Now().In(c.tz).Add(time.Minute * time.Duration(c.cfg.JWT.JwtExpirationTime)).Unix()
	accessToken, vErr := c.generateTokenClaim(int64(uid), accessTokenTime)
	if err != nil {
		return nil, vErr
	}

	return &ACToken{AccessToken: accessToken, AccessTokenExpire: accessTokenTime}, nil
}

func (c *defaultClient) generateTokenClaim(uid int64, exp int64) (string, errors.Error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = c.cfg.JWT.JwtIssuer
	claims["sub"] = fmt.Sprintf("%d", uid)
	claims["exp"] = exp
	claims["iat"] = time.Now().In(c.tz).Unix()

	secretKey := []byte(c.cfg.JWT.JwtSecret)
	t, err := token.SignedString(secretKey)
	if err != nil {
		return "", errors.ErrClientTokenInvalid
	}

	return t, nil
}
