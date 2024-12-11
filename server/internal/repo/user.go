package repo

import (
	"context"
	"knowFood/internal/model"
)

// UserRepo 用户repo接口
type UserRepo interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserById(ctx context.Context, uid int64) (*model.User, error)
	GetUserByMobile(ctx context.Context, mobile string) (*model.User, error)
	GetUserByOpenID(ctx context.Context, openID string) (*model.User, error)
	CreateUser(ctx context.Context, userInfo *model.User) error
}
