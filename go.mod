module orm-demo

go 1.12

require (
	github.com/gogo/protobuf v1.3.0
	github.com/golang/protobuf v1.3.2
	github.com/google/wire v0.3.0
	github.com/smartystreets/goconvey v1.6.4
	go-common v1.26.0
	go.uber.org/automaxprocs v1.4.0
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	google.golang.org/grpc v1.27.1
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.7
)

replace go-common => git.bilibili.co/platform/go-common v1.26.0
