// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description 在这里像外部提供wire工具使用的ProviderSet

package mysql

import (
	"knowFood/internal/repo"
	"knowFood/utils/db"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	// 事务对象
	NewTransaction,
	wire.Bind(new(db.Transaction), new(*transaction)),
	// 用户的repo层
	NewUserRepo,
	wire.Bind(new(repo.UserRepo), new(*userRepo)),
)
