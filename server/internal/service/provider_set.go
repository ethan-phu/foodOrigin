// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description service层ProviderSet

package service

import "github.com/google/wire"

// ProviderSet is a wire provider set for services
var ProviderSet = wire.NewSet(
	NewUserService,
	wire.Bind(new(UserService), new(*userService)),
	NewTxDemoService,
	wire.Bind(new(TxDemoService), new(*txDemoService)),
)
