package services

import (
	"fmt"
	"github.com/kz-login/pkg/errors"
)

func (s *defaultService) Logout(userId int64) errors.Error {
	key := fmt.Sprintf("auth#%d", userId)
	isDel, err := s.rdc.Del(key)
	if err != nil {
		return errors.NewError(errors.ErrCodeInternalServer, err.Error())
	}

	if isDel == 0 {
		return errors.NewError(errors.ErrCodeInternalServer, "key not found")
	}

	return nil
}
