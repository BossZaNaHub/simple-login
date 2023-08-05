package services

import (
	"fmt"
	"github.com/kz-login/app/models"
	"github.com/kz-login/pkg/encrypt"
	"github.com/kz-login/pkg/errors"
	"github.com/kz-login/pkg/jwt"
)

func (s *defaultService) Login(data *models.MemberLoginData) (*models.MemberData, *models.JwtToken, errors.Error) {
	user, err := s.rp.GetUserByMobileNumber(data.MobileNumber)
	if err != nil {
		return nil, nil, errors.ErrClientMemberNotFound
	}

	if !user.IsActive {
		return nil, nil, errors.ErrClientMemberNotVerified
	}

	encryptPwd := encrypt.MD5(data.Password)
	if user.PasswordEncrypted != encryptPwd {
		return nil, nil, errors.ErrClientPasswordMismatch
	}

	/* JWT Encrypted */
	ac, err := s.csJwt.CreateToken(jwt.ClaimTokenData{
		UID:          user.Id,
		Name:         user.Name,
		MobileNumber: user.MobileNumber,
	})

	/* Store to redis */
	key := fmt.Sprintf("auth#%d", user.Id)
	sErr := s.rdc.Set(key, ac.AccessToken)
	if err != nil {
		return nil, nil, errors.NewError(errors.ErrCodeInternalServer, sErr.Error())
	}

	return user, &models.JwtToken{
		AccessToken:         ac.AccessToken,
		AccessTokenExpired:  ac.AccessTokenExpire,
		RefreshToken:        ac.RefreshToken,
		RefreshTokenExpired: ac.RefreshTokenExpire,
	}, nil
}
