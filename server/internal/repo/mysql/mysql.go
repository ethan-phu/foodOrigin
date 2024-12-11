// author: xmgtony
// date: 2023-06-29 14:47
// version:

package mysql

import (
	"context"
	"knowFood/internal/model"
	"knowFood/utils/config"
	"knowFood/utils/db"

	"gorm.io/gorm"
)

// var _ IDataSource = new(*defaultMysqlDataSource) 也可
var _ db.IDataSource = (*defaultMysqlDataSource)(nil)

// defaultMysqlDataSource 默认mysql数据源实现
type defaultMysqlDataSource struct {
	master *gorm.DB // 定义私有属性，用来持有主库链接，防止每次创建，创建后直接返回该变量。
	slave  *gorm.DB // 同上，从库链接
}

func (d *defaultMysqlDataSource) Master(ctx context.Context) *gorm.DB {
	// 事物, 根据事物的key取出tx，value是对应的db数据库连接对象
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	if d.master == nil {
		panic("The [master] connection is nil, Please initialize it first.")
	}
	return d.master
}

func (d *defaultMysqlDataSource) Slave(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	if d.slave == nil {
		panic("The [slave] connection is nil, Please initialize it first.")
	}
	return d.slave
}

func (d *defaultMysqlDataSource) Close() {
	// 关闭主库链接
	if d.master != nil {
		m, err := d.master.DB()
		if err != nil {
			_ = m.Close()
		}
	}
	// 关闭从库链接
	if d.slave != nil {
		s, err := d.slave.DB()
		if err != nil {
			_ = s.Close()
		}
	}
}

func (d *defaultMysqlDataSource) Migrate() {
	// 主库的数据库迁移
	if d.master != nil {
		migrage(d.master)
	}
	// 从库迁移
	if d.slave != nil {
		migrage(d.slave)
	}
}

func migrage(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}

func NewDefaultMysql(c config.DBConfig) *defaultMysqlDataSource {
	var dbOrm *gorm.DB
	if c.Sqlite {
		dbOrm = db.GetSqliteConn(c.MaximumPoolSize, c.MaximumIdleSize)
	} else {
		dbOrm = db.GetMysqlConn(
			c.Username,
			c.Password,
			c.Host,
			c.Port,
			c.Dbname,
			c.MaximumPoolSize,
			c.MaximumIdleSize)
	}
	return &defaultMysqlDataSource{
		master: dbOrm,
	}
}
