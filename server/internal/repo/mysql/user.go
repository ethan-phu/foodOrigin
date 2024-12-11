package mysql

import (
	"context"
	"knowFood/internal/model"
	"knowFood/internal/repo"
	"knowFood/utils/db"
)

var _ repo.UserRepo = (*userRepo)(nil)

type userRepo struct {
	ds db.IDataSource
}

func NewUserRepo(_ds db.IDataSource) *userRepo {
	return &userRepo{
		ds: _ds,
	}
}

func (ur *userRepo) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user := &model.User{}
	err := ur.ds.Master(ctx).Where("name = ?", name).Find(user).Error
	return user, err
}

func (ur *userRepo) GetUserById(ctx context.Context, uid int64) (*model.User, error) {
	user := &model.User{}
	err := ur.ds.Master(ctx).Where("id = ?", uid).First(user).Error
	return user, err
}

func (ur *userRepo) GetUserByMobile(ctx context.Context, mobile string) (*model.User, error) {
	user := &model.User{}
	err := ur.ds.Master(ctx).
		Where("mobile = ?", mobile).
		Where("enabled_status = 1").
		First(user).Error
	return user, err
}

func (ur *userRepo) CreateUser(ctx context.Context, userInfo *model.User) error {
	return ur.ds.Master(ctx).Create(userInfo).Error
}

// GetUserByOpenID 通过OpenID获取用户信息
func (ur *userRepo) GetUserByOpenID(ctx context.Context, openID string) (*model.User, error) {
	user := &model.User{}
	err := ur.ds.Master(ctx).
		Where("open_id = ?", openID).
		Where("enabled_status = 1").
		First(user).Error
	return user, err
}
