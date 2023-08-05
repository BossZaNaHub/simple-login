package services

import (
	"fmt"
	"github.com/kz-login/app/models"
	"github.com/kz-login/pkg/errors"
)

func (s *defaultService) Refresh(userId int64, rfToken string) (*models.JwtToken, errors.Error) {
	key := fmt.Sprintf("auth#%d", userId)

	token, rErr := s.rdc.Get(key)
	if rErr != nil || token == "" {
		return nil, errors.ErrClientUnauthorized
	}

	t, err := s.csJwt.RefreshToken(rfToken)
	if err != nil {
		return nil, err
	}

	return &models.JwtToken{AccessToken: t.AccessToken}, nil
}
