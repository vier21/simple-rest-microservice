package service

import (
	"context"
	"fmt"
	"gorest/internal/utils"
	"gorest/internal/utils/errors"
	"gorest/internal/utils/hash"
	"gorest/pkg/user/domain/web"
	"gorest/pkg/user/repository"
	"slices"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type service struct {
	repo      repository.UserRepository
	validator validator.Validate
}

var (
	ErrRequestNotValid = "Input not valid: %s"
	ErrResourceFound   = "resource not found: %s"
)

func NewService(userRepo repository.UserRepository) *service {
	return &service{
		repo: userRepo,
	}
}

func (s *service) RegisterUser(ctx context.Context, req web.RegisterRequest) *utils.Result {
	err := utils.ValidateStruct(&s.validator, req)

	if err != nil {
		return &utils.Result{
			Data:  nil,
			Error: errors.BadRequest(ErrRequestNotValid),
		}
	}

	convUser := web.ToUserDomain(req)
	usrMatch, _ := s.repo.GetUserByUsername(ctx, convUser.Username)

	if usrMatch.Username == convUser.Username {
		return &utils.Result{
			Data:  nil,
			Error: errors.BadRequest("username already exist"),
		}
	}

	if usrMatch.Email == convUser.Email {
		return &utils.Result{
			Data:  nil,
			Error: errors.BadRequest("email already exist"),
		}
	}

	hashPassword := hash.GenerateHashPassword(convUser.Password)
	convUser.Id = uuid.New()
	convUser.Password = hashPassword

	user, err := s.repo.InsertOneUser(ctx, convUser)

	if err != nil {
		return &utils.Result{
			Data:  nil,
			Error: errors.InternalServerError("error occur when user register"),
		}
	}

	userResp := web.ToUserResponse(*user)

	return &utils.Result{
		Data:  userResp,
		Error: nil,
	}
}

func (s *service) BulkRegisterUser(ctx context.Context, req web.RegisterRequests) *utils.Result {
	for _, v := range req {
		if err := utils.ValidateStruct(&s.validator, v); err != nil {
			return &utils.Result{
				Data:  nil,
				Error: errors.BadRequest(fmt.Sprintf("one of user data not valid: %s", err.Error())),
			}
		}
	}

	userIn := web.ToUserDomains(req)
	for _, v := range userIn {
		v.Id = uuid.New()
		v.Password = hash.GenerateHashPassword(v.Password)
	}

	users, err := s.repo.InsertManyUser(ctx, userIn)

	if err != nil {
		return &utils.Result{
			Data:  nil,
			Error: errors.BadRequest(fmt.Sprintf("error perform bulk register: %s", err.Error())),
		}
	}

	data := map[string]any{
		"inserted_count":     0,
		"not_inserted_count": 0,
		"user_inserted":      []string{},
		"user_not_inserted":  []string{},
	}

	for i, v := range users {
		contains := slices.Contains(userIn, v)

		if contains {
			data["inserted_count"] = data["inserted_count"].(int) + 1
			data["user_inserted"] = append(data["user_inserted"].([]string), userIn[i].Username)
		} else {
			data["not_inserted_count"] = data["not_inserted_count"].(int) + 1
			data["user_not_inserted"] = append(data["user_not_inserted"].([]string), userIn[i].Username)
		}
	}

	return &utils.Result{
		Data:  data,
		Error: nil,
	}
}

func (s *service) EditUser(ctx context.Context, req web.UpdateRequest) *utils.Result {

	return nil
}

func (s *service) GetAllUser(ctx context.Context) *utils.Result {
	return nil
}

func (s *service) GetUserById(ctx context.Context, id string) (*web.UserResponse, error) {
	return nil, nil
}

// can be use uuid.FromString(id)
func (s *service) GetUserByUsername(ctx context.Context, username string) (*web.UserResponse, error) {
	return nil, nil
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (*web.UserResponse, error) {
	return nil, nil
}

func (s *service) DeleteUser(ctx context.Context, id string) (string, error) {
	return "", nil
}

func (s *service) DeleteManyUser(ctx context.Context, ids []string) (string, error) {
	return "", nil

}
