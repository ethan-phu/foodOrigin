// Created on 2021/5/4.
// @author tony
// email xmgtony@gmail.com
// description 配置mysql链接

package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// IDataSource 定义数据库数据源接口，按照业务需求可以返回主库链接Master和从库链接Slave
type IDataSource interface {
	Master(ctx context.Context) *gorm.DB
	Slave(ctx context.Context) *gorm.DB
	Close()
}

// GetSqliteConn 创建SQLite链接
func GetSqliteConn(maxPoolSize, maxIdle int) *gorm.DB {
	dbPath := "./data_base.db"
	// 检查数据库文件是否存在
	_, err := os.Stat(dbPath)
	if err != nil && os.IsNotExist(err) {
		// 如果文件不存在，则创建文件
		f, err := os.Create(dbPath)
		if err != nil {
			log.Fatalf("Error creating SQLite database file: %v", err)
		}
		defer f.Close()
	}
	// SQLite 数据库文件路径
	dsn := dbPath
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		// 其他配置项可以保留，比如PrepareStmt
		PrepareStmt: true, // 缓存每一条sql语句，提高执行速度
	})
	if err != nil {
		panic(err)
	}
	// 获得底层的sql.DB对象
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	// 设置连接池参数
	sqlDb.SetMaxOpenConns(maxPoolSize)
	sqlDb.SetMaxIdleConns(maxIdle)
	// 设置连接的最大生命周期
	sqlDb.SetConnMaxLifetime(time.Hour)
	return db
}

// GetMysqlConn 创建Mysql链接
func GetMysqlConn(user, password, host, port, dbname string, maxPoolSize, maxIdle int) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true, // 缓存每一条sql语句，提高执行速度
	})
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetConnMaxLifetime(time.Hour)
	// 设置连接池大小
	sqlDb.SetMaxOpenConns(maxPoolSize)
	sqlDb.SetMaxIdleConns(maxIdle)
	return db
}
