// Created on 2023/3/15.
// @author tony
// email xmgtony@gmail.com
// description 事物控制接口

package mysql

import (
	"context"
	"knowFood/utils/db"

	"gorm.io/gorm"
)

type contextTxKey struct{}

// 事物默认实现
type transaction struct {
	ds db.IDataSource
}

func NewTransaction(_ds db.IDataSource) *transaction {
	return &transaction{ds: _ds}
}

// 事务接口，执行事务的函数execute调用
func (t *transaction) Execute(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.ds.Master(ctx).Transaction(func(tx *gorm.DB) error {
		/*
			context.WithValue 函数是 Go 语言 context 包中提供的一个函数，它用于创建一个新的 context 对象，这个新对象与原始的 context 对象有相同的父级（即相同的取消信号和超时设置），但带有附加的值。

			这里是 WithValue 函数的基本用法和作用：

			创建新的 context 对象：WithValue 函数接收三个参数：一个 context.Context 对象，一个键（通常是实现了 fmt.Stringer 接口的类型），和一个值。它返回一个新的 context.Context 对象。

			存储键值对：这个新 context 对象将存储传入的键值对，可以通过调用 Context.Value 方法来检索这个值。

			线程安全：context 是线程安全的，可以在多个 goroutine 中安全使用。

			使用 WithValue 的目的通常是为了在不同的函数调用之间传递额外的信息，而不需要通过参数显式传递。这在处理例如数据库事务或请求 ID 等跨越多个层的请求处理时非常有用。
		*/
		// 创建一个新context，并往里面存储了一个tx，key是contextTxKey{}
		withValue := context.WithValue(ctx, contextTxKey{}, tx)
		return fn(withValue)
	})
}
