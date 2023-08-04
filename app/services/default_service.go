package services

import (
	"github.com/kz-login/app/models"
	"github.com/kz-login/app/repositories"
	"github.com/kz-login/pkg/errors"
	customjwt "github.com/kz-login/pkg/jwt"
)

type defaultService struct {
	rp    repositories.Repository
	csJwt customjwt.Client
}

type Service interface {
	Login(data *models.MemberLoginData) (*models.MemberData, *models.JwtToken, errors.Error)
	Profile(userId int64) (*models.MemberData, errors.Error)
}

func NewService(rp repositories.Repository, csJwt customjwt.Client) Service {
	return &defaultService{rp: rp, csJwt: csJwt}
}
