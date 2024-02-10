package web

import (
	"gorest/pkg/user/domain"
	"time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Id        uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"userName" validate:"required"`
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Email     string    `json:"email" validate:"email,required"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Role      string    `json:"role,omitempty"`
}

type RegisterRequests []RegisterRequest 

func ToUserDomains(regs RegisterRequests) []domain.User {
	uDomains := make([]domain.User, len(regs))
	for i := range regs {
		uDomains[i].Id = regs[i].Id
		uDomains[i].Username = regs[i].Username
		uDomains[i].FirstName = regs[i].FirstName
		uDomains[i].LastName = regs[i].LastName
		uDomains[i].Email = regs[i].Email
		uDomains[i].Password = regs[i].Password
		uDomains[i].CreatedAt = regs[i].CreatedAt
		uDomains[i].Role = regs[i].Role
	}

	return uDomains
}

func ToUserDomain(reg RegisterRequest) domain.User {
	return domain.User{
		Id:        reg.Id,
		Username:  reg.Username,
		FirstName: reg.FirstName,
		LastName:  reg.LastName,
		Email:     reg.Email,
		Password:  reg.Password,
		CreatedAt: reg.CreatedAt,
		Role:      reg.Role,
	}
}

type UserResponse struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"userName"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
	}
}

func ToUserResponses(users []domain.User) UserResponses {
	uResponses := make(UserResponses, len(users))
	for k, v := range users {
		uResponses[k].Id = v.Id
		uResponses[k].Username = v.Username
		uResponses[k].FirstName = v.FirstName
		uResponses[k].LastName = v.LastName
		uResponses[k].Email = v.Email
		uResponses[k].Role = v.Role
	}

	return uResponses
}

type UpdateRequest struct {
	Username  string `json:"userName,omitempty" validate:"required"`
	FirstName string `json:"firstName,omitempty" validate:"required"`
	LastName  string `json:"lastName,omitempty" validate:"required"`
	Email     string `json:"email,omitempty" validate:"email,required"`
	Password  string `json:"password,omitempty" validate:"required"`
	Role      string `json:"role,omitempty"`
}

func ToUpdateDomain(upd UpdateRequest) domain.User {
	return domain.User{
		Username:  upd.Username,
		FirstName: upd.FirstName,
		LastName:  upd.LastName,
		Email:     upd.Email,
		Password:  upd.Password,
		Role:      upd.Role,
	}
}


type ResponsePayload struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type UserResponses []UserResponse
