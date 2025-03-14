package dto

import (
	"time"
	"github.com/google/uuid"
	"fusossafuoye.ng/app/model"
)

type UserResponse struct {
	ID        uuid.UUID     `json:"id"`
	Name  string     `json:"name"`
	
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func ToUserResponse(user *model.UserModel) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserListResponse(users []model.UserModel) []UserResponse {
	response := make([]UserResponse, 0)
	for _, user := range users {
		response = append(response, *ToUserResponse(&user))
	}
	return response
}
