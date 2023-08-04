package repositories

import (
	"github.com/kz-login/app/models"
	"github.com/kz-login/pkg/db"
	"github.com/kz-login/pkg/errors"
)

type defaultRepository struct {
	client db.Client
}

type Repository interface {
	GetUserByMobileNumber(mobileNumber string) (*models.MemberData, errors.Error)
	GetUserById(id int64) (*models.MemberData, errors.Error)
}

func NewRepository(client db.Client) Repository {
	return &defaultRepository{
		client: client,
	}
}
