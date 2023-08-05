package jwt

import (
	"github.com/kz-login/pkg/errors"
)

type Client interface {
	CreateToken(data ClaimTokenData) (*ACToken, errors.Error)
	RefreshToken(rfToken string) (*ACToken, errors.Error)
}
