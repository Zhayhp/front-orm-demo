// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"orm-demo/internal/dao"
	"orm-demo/internal/service"
	"orm-demo/internal/server/grpc"
	"orm-demo/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
