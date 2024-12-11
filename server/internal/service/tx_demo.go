// author: xmgtony
// date: 2023-06-29 15:00
// version: 事务操作演示

package service

import (
	"context"
	"knowFood/utils/db"
)

// TxDemoService txDemo服务接口
type TxDemoService interface {
	SaveWithTx(ctx context.Context)
}

// txDemoService 默认实现
type txDemoService struct {
	userService UserService
	tx          db.Transaction
}

// NewTxDemoService creates a new instance of txDemoService
func NewTxDemoService(us UserService, tx db.Transaction) *txDemoService {
	return &txDemoService{
		userService: us,
		tx:          tx,
	}
}

// SaveWithTx demonstrates transaction usage
func (s *txDemoService) SaveWithTx(ctx context.Context) {
	err := s.tx.Execute(ctx, func(ctx context.Context) error {
		// Add your transaction logic here
		return nil
	})
	if err != nil {
		// Handle error
		return
	}
}
