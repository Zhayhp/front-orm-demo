package dao

import (
	"context"
	"gorm.io/gorm"
	"time"

	"go-common/library/cache/memcache"
	"go-common/library/cache/redis"
	"go-common/library/conf/paladin"
	"go-common/library/sync/pipeline/fanout"
	xtime "go-common/library/time"
	"orm-demo/internal/model"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB)

//go:generate kratos tool btsgen
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	Article(c context.Context, id int64) (*model.Article, error)
	DB() *gorm.DB
}

// dao dao.
type dao struct {
	db         *gorm.DB
	redis      *redis.Redis
	mc         *memcache.Memcache
	cache      *fanout.Fanout
	demoExpire int32
}

// New new a dao and return.
func New(db *gorm.DB) (d Dao, cf func(), err error) {
	return newDao(db)
}

func newDao(db *gorm.DB) (d *dao, cf func(), err error) {
	var cfg struct {
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &dao{
		db:         db,
		cache:      fanout.New("cache"),
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.cache.Close()
}

// Close close the resource.
func (d *dao) DB() *gorm.DB {
	return d.db
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
