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

func NewService(userRepo repository.UserRepository, val *validator.Validate) *service {
	return &service{
		repo:      userRepo,
		validator: *val,
	}
}

func (s *service) RegisterUser(ctx context.Context, req web.RegisterRequest) utils.Result {
	err := s.validator.Struct(req)

	if err != nil {
		return utils.Result{
			Data:  nil,
			Error: errors.BadRequest(fmt.Sprintf(ErrRequestNotValid, err.Error())),
		}
	}

	convUser := web.ToUserDomain(req)
	usrMatch, _ := s.repo.GetUserByUsername(ctx, convUser.Username)

	if usrMatch != nil {
		if usrMatch.Username == convUser.Username {
			return utils.Result{
				Data:  nil,
				Error: errors.BadRequest("username already exist"),
			}
		}

		if usrMatch.Email == convUser.Email {
			return utils.Result{
				Data:  nil,
				Error: errors.BadRequest("email already exist"),
			}
		}
	}

	hashPassword := hash.GenerateHashPassword(convUser.Password)
	convUser.Id = uuid.New()
	convUser.Password = hashPassword

	if convUser.Role == "" {
		convUser.Role = "user"
	}

	user, err := s.repo.InsertOneUser(ctx, convUser)

	if err != nil {
		return utils.Result{
			Data:  nil,
			Error: errors.InternalServerError("error occur when user register"),
		}
	}

	userResp := web.ToUserResponse(*user)

	return utils.Result{
		Data:  userResp,
		Error: nil,
	}
}

func (s *service) BulkRegisterUser(ctx context.Context, req web.RegisterRequests) utils.Result {
	for i := range req {
		if err := s.validator.Struct(req[i]); err != nil {
			return utils.Result{
				Data:  nil,
				Error: errors.BadRequest(ErrRequestNotValid),
			}
		}
	}

	userIn := web.ToUserDomains(req)
	for i := range userIn {
		userIn[i].Id = uuid.New()
		userIn[i].Password = hash.GenerateHashPassword(userIn[i].Password)
	}

	users, err := s.repo.InsertManyUser(ctx, userIn)
	fmt.Println(users)
	if err != nil {
		return utils.Result{
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

	for i := range users {
		contains := slices.Contains(userIn, userIn[i])

		if contains {
			data["inserted_count"] = data["inserted_count"].(int) + 1
			data["user_inserted"] = append(data["user_inserted"].([]string), userIn[i].Username)
		} else {
			data["not_inserted_count"] = data["not_inserted_count"].(int) + 1
			data["user_not_inserted"] = append(data["user_not_inserted"].([]string), userIn[i].Username)
		}
	}

	return utils.Result{
		Data:  data,
		Error: nil,
	}
}

func (s *service) EditUser(ctx context.Context, id string, req web.UpdateRequest) utils.Result {
	uid, err := uuid.Parse(id)

	if err != nil {
		return utils.Result{
			Data:  nil,
			Error: errors.InternalServerError("failed to parse uuid"),
		}
	}

	user, _ := s.repo.GetUserById(ctx, uid)

	if user == nil {
		return utils.Result{
			Data:  nil,
			Error: errors.NotFound("user with given id not found"),
		}
	}

	updateIn := web.ToUpdateDomain(req)
	out, err := s.repo.UpdateOneUser(ctx, updateIn)
	if err != nil {
		return utils.Result{
			Data:  nil,
			Error: errors.InternalServerError("failed to parse uuid"),
		}
	}
	return utils.Result{
		Data:  out,
		Error: nil,
	}
}

func (s *service) GetAllUser(ctx context.Context) utils.Result {
	users, err := s.repo.GetAllUser(ctx)
	if err != nil {
		return utils.Result{
			Data:  nil,
			Error: errors.InternalServerError("cannot fetch users"),
		}
	}

	return utils.Result{
		Data:  users,
		Error: nil,
	}
}

func (s *service) GetUserById(ctx context.Context, id string) utils.Result {
	uid, err := uuid.Parse(id)

	if err != nil {
		return utils.Result{
			Data:  nil,
			Error: errors.BadRequest("invalid id format"),
		}
	}

	user, err := s.repo.GetUserById(ctx, uid)

	if err != nil {
		return utils.Result{
			Data:  nil,
			Error: errors.BadRequest("id not found"),
		}
	}
	return utils.Result{
		Data:  user,
		Error: nil,
	}
}

// can be use uuid.FromString(id)
func (s *service) GetUserByUsername(ctx context.Context, username string) utils.Result {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return utils.Result{
			Data:  nil,
			Error: errors.BadRequest("user not found invalid username"),
		}
	}

	return utils.Result{
		Data:  user,
		Error: nil,
	}
}

func (s *service) GetUserByEmail(ctx context.Context, email string) utils.Result {
	user, err := s.repo.GetUserByUsername(ctx, email)
	if err != nil {
		return utils.Result{
			Data:  nil,
			Error: errors.BadRequest("user not found invalid email"),
		}
	}

	return utils.Result{
		Data:  user,
		Error: nil,
	}
}

func (s *service) DeleteUser(ctx context.Context, id string) utils.Result {
	uid, err := uuid.Parse(id)

	if err != nil {
		return utils.Result{
			Data:  nil,
			Error: errors.BadRequest("invalid id format"),
		}
	}

	user, _ := s.repo.GetUserById(ctx, uid)

	if user == nil {
		return utils.Result{
			Data:  nil,
			Error: errors.NotFound("user wich given id not found"),
		}
	}

	del, err := s.repo.DeleteOneUser(ctx, *user)
	if err != nil {
		return utils.Result{
			Data:  nil,
			Error: errors.InternalServerError("error when perform delete user"),
		}
	}
	return utils.Result{
		Data:  del,
		Error: nil,
	}
}

func (s *service) DeleteManyUser(ctx context.Context, ids []string) utils.Result {
	return utils.Result{
		Data:  nil,
		Error: nil,
	}
}
