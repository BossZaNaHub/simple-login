package services

import (
	"fmt"
	"github.com/kz-login/app/models"
	"github.com/kz-login/pkg/errors"
)

func (s *defaultService) Profile(userId int64) (*models.MemberData, errors.Error) {
	key := fmt.Sprintf("auth#%d", userId)
	token, rErr := s.rdc.Get(key)
	if rErr != nil || token == "" {
		return nil, errors.ErrClientUnauthorized
	}
	user, err := s.rp.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	if !user.IsActive {
		return nil, errors.ErrClientMemberNotVerified
	}

	return user, nil
}
