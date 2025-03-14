package request

import (
	"fusossafuoye.ng/app/dao"
	"fusossafuoye.ng/app/model"
)

type RegisterUserRequestData struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *RegisterUserRequestData) ToModel() *model.UserModel {
	return &model.UserModel{
		User: dao.User{
			FullName: r.FullName,
			Email:    r.Email,
			Password: r.Password,
		},
	}
}
