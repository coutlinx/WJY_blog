/*
@Time : 2024/12/6 14:43
@Author : linx
@File : models.go
@dsc: 公用的模型类
*/

package models

import (
	"blog/configs"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"time"
)

type BaseModel struct {
	Id         int64     `gorm:"primaryKey;AUTO_INCREMENT;column:Id;type:bigint(20);comment:Id;NOT NULL" json:"id"`
	CreateTime time.Time `gorm:"index;autoCreateTime;column:CreateTime;type:datetime;comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"autoUpdateTime;column:UpdateTime;type:datetime;comment:更新时间" json:"update_time"`
}

func AutoMigrate() error {
	db := configs.Config.DB
	// 获取现有的日志器配置
	newLogger := db.Logger

	// 设置为无日志输出
	db.Logger = db.Logger.LogMode(logger.Silent)
	err := db.AutoMigrate(
		&BlogUser{},
	)
	if err != nil {
		return err
	}

	db.Logger = newLogger
	zap.L().Info("All tables AutoMigrate success! ")
	return nil
}
