package services

import (
	"github.com/kz-login/app/models"
	"github.com/kz-login/pkg/errors"
)

func (s *defaultService) Refresh(rfToken string) (*models.JwtToken, errors.Error) {
	t, err := s.csJwt.RefreshToken(rfToken)
	if err != nil {
		return nil, err
	}

	return &models.JwtToken{AccessToken: t.AccessToken}, nil
}
