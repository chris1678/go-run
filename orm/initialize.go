package orm

import (
	"fmt"
	"github.com/chris1678/go-run/config"
	"github.com/chris1678/go-run/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var _db *gorm.DB

func Db() *gorm.DB {
	return _db
}
func Initialize() {
	c := config.DatabaseConfig

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%dms",
		c.User,
		c.Password,
		c.Addr,
		c.Port,
		c.DbName,
		c.Timeout,
	)
	var err error
	_db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: true, //是否跳过事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //屏蔽复数表名
		},
		Logger: logger.New(
			gormlog.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel:      3,
			},
		),
		DisableForeignKeyConstraintWhenMigrating: true, //创建表的是否取消外键约束，不使用数据库自己的外键约束
	})

	if err != nil {
		logger.LogHelper.Fatal(err)
	}
	sqlDB, err2 := _db.DB()
	if err2 != nil {
		logger.LogHelper.Fatal(err2)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(c.ConnMaxLifeTime) * time.Hour)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime) * time.Hour)

}
