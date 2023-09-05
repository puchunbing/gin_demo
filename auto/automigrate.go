package auto

import (
	"gin_test/config"
	"gin_test/model"
)

// 数据库迁移
func AutoMigrate() {
	db := config.GetMySql()
	db.AutoMigrate(new(model.Info))
}
