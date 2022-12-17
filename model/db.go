package model

import (
	"DineEasy/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Model struct {
	ID        uint      `gorm:"comment:主键ID;primarykey;" json:"id"`
	CreatedAt time.Time `gorm:"comment:创建时间;" json:"created_at;"`
	UpdatedAt time.Time `gorm:"comment:更新时间;" json:"updated_at;"`
}

var (
	Db  *gorm.DB
	err error
)

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPost, utils.DbName)
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,  // string 类型字段的默认长度
		DisableDatetimePrecision:  true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: true, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败,%s", err))
	}
	sqlDB, err := Db.DB()
	if err != nil {
		panic(fmt.Sprintf("数据库初始化配置失败,%s", err))
	}
	err = Db.AutoMigrate(&Store{}, &Member{}, &Category{}, &Item{}, &Combo{}, &ComboGroup{}, &ComboGroupItem{}, &Cart{}, &Table{}, &Order{})

	if err != nil {
		panic(fmt.Sprintf("数据库迁移失败,%s", err))
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
