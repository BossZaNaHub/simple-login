package repositories

import (
	"fmt"
	"github.com/kz-login/app/models"
	"github.com/kz-login/pkg/errors"
)

func (r *defaultRepository) GetUserById(id int64) (*models.MemberData, errors.Error) {
	user, err := r.client.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return &models.MemberData{
		Id:                int64(user.ID),
		Name:              fmt.Sprintf("%v %v", user.Firstname, user.Lastname),
		Email:             user.Email,
		MobileNumber:      user.MobileNumber,
		BirthOfDate:       user.Birthday,
		PasswordEncrypted: user.PasswordEncrypted,
		IsActive:          user.IsActive,
	}, nil
}
