package main

import (
	consul "github.com/asim/go-micro/plugins/registry/consul/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/opentracing/opentracing-go"
	"github.com/tupig-7/cart/common"
	"github.com/tupig-7/cart/domain/repository"
	service2 "github.com/tupig-7/cart/domain/service"
	"github.com/tupig-7/cart/handler"
	pb "github.com/tupig-7/cart/proto"

	"github.com/micro/micro/v3/service/logger"
)
var QPS = 100
func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("106.14.89.192", 8500, "/micro/config")
	if err != nil {
		logger.Error(err)
	}
	//注册中心
	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"106.14.89.192:8500",
		}
	})
	//链路追踪
	t, io, err := common.NewTracer("go.micro.service.cart", "106.14.89.192:6831")
	if err != nil {
		logger.Error(err)
	}

	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//数据库连接
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+":"+mysqlInfo.Port+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)
	err = repository.NewCartRepository(db).InitTable()
	if err != nil {
		logger.Error(err)
	}
	// Create service
	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8081"),
		micro.Registry(consul),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		)
	service.Init()

	cartDataService := service2.NewCartDataService(repository.NewCartRepository(db))
	// Register handler
	pb.RegisterCartHandler(service.Server(), &handler.Cart{CartDataService: cartDataService})

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
