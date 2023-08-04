package db

import (
	"github.com/kz-login/pkg/db/daos"
	"github.com/kz-login/pkg/errors"
)

type Client interface {
	AutoMigrate()
	Seed()

	GetUserById(userId int64) (*daos.User, errors.Error)
	GetUserByMobileNumber(mobileNumber string) (*daos.User, errors.Error)
}
