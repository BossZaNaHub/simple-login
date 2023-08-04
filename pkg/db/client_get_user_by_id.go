package db

import (
	"github.com/kz-login/pkg/db/daos"
	"github.com/kz-login/pkg/errors"
)

func (c *defaultClient) GetUserById(userId int64) (*daos.User, errors.Error) {
	var user daos.User
	err := c.Conn.First(&user, userId).Error
	if err != nil {
		return nil, errors.NewDefaultError(err)
	}

	return &user, nil
}
