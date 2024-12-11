// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description user handler层ProviderSet

package v1

import (
	"knowFood/internal/handler/v1/user"

	"github.com/google/wire"
)

// handler的提供者
/*
用户的handler对象
账号账单的handler对象
*/
var ProviderSet = wire.NewSet(
	user.NewUserHandler,
)
