package dao

import (
	"context"
	"go-common/library/conf/paladin"
	"go-common/library/database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"orm-demo/internal/model"

	"gorm.io/gorm"
)

func NewDB() (db *gorm.DB, cf func(), err error) {
	var (
		cfg sql.Config
		ct  paladin.TOML
	)
	if err = paladin.Get("db.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	if db, err = gorm.Open(mysql.Open(cfg.DSN), nil); err != nil {
		return
	}
	db.Logger.LogMode(logger.Info)
	cf = func() {}
	return
}

func (d *dao) RawArticle(ctx context.Context, id int64) (art *model.Article, err error) {
	// get data from db
	return
}
