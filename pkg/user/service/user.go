package service

import (
	"context"
	"gorest/internal/utils"
	"gorest/pkg/user/domain/web"
)

type UserService interface {
	RegisterUser(ctx context.Context, req web.RegisterRequest) *utils.Result
	BulkRegisterUser(ctx context.Context, req web.RegisterRequests) *utils.Result
	EditUser(ctx context.Context, req web.UpdateRequest) *utils.Result
	GetAllUser(ctx context.Context) *utils.Result
	GetUserById(ctx context.Context, id string) *utils.Result //can be use uuid.FromString(id)
	GetUserByUsername(ctx context.Context, username string) *utils.Result
	GetUserByEmail(ctx context.Context, email string) *utils.Result
	DeleteUser(ctx context.Context, id string) *utils.Result
	DeleteManyUser(ctx context.Context, ids []string) *utils.Result
}
