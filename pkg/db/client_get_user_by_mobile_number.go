package db

import (
	"github.com/kz-login/pkg/db/daos"
	"github.com/kz-login/pkg/errors"
)

func (c *defaultClient) GetUserByMobileNumber(mobile string) (*daos.User, errors.Error) {
	var user daos.User
	err := c.Conn.Where("mobile_number = ?", mobile).First(&user).Error
	if err != nil {
		return nil, errors.NewDefaultError(err)
	}

	return &user, nil
}
