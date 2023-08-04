package services

import (
	"github.com/kz-login/app/models"
	"github.com/kz-login/pkg/errors"
)

func (s *defaultService) Profile(userId int64) (*models.MemberData, errors.Error) {
	user, err := s.rp.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	if !user.IsActive {
		return nil, errors.ErrClientMemberNotVerified
	}

	return user, nil
}
