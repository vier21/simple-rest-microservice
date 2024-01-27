package web

import (
	"gorest/pkg/user/domain"
	"time"
)

type RegisterRequest struct {
	Username  string    `json:"userName" validate:"required"`
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Email     string    `json:"emial" validate:"email,required"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Role      string    `json:"role,omitempty"`
}

type UserResponse struct {
	Username  string    `json:"userName" validate:"required"`
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Email     string    `json:"emial" validate:"email,required"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Role      string    `json:"role,omitempty"`
}

func ToUserResponse(user *domain.User) *UserResponse {
	return &UserResponse{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Role:      user.Role,
	}
}
