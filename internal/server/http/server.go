package http

import (
	"encoding/json"
	"go-common/library/conf/paladin"
	"go-common/library/ecode"
	"go-common/library/log"
	bm "go-common/library/net/http/blademaster"
	"io/ioutil"
	"net/http"
	pb "orm-demo/api"
	"orm-demo/internal/model"
	"orm-demo/internal/service"
)

var svc *service.Service

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	engine = bm.DefaultServer(&cfg)
	pb.RegisterDemoBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/orm-demo")
	{
		g.GET("/gorm", howToStart)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	var req = new(model.RetrieveReq)
	bs, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("ioutil.ReadAll() error(%v)", err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.Request.Body.Close()
	err = json.Unmarshal(bs, req)
	if err != nil {
		log.Error("howToStart error(%+v)", err)
		c.JSON(req, err)
		return
	}
	log.Info("req: %+v", req)
	c.JSON(svc.Retrieve(c, req))
}
