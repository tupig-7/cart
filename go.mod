module github.com/tupig-7/cart

go 1.15

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/asim/go-micro/plugins/config/source/consul/v3 v3.0.0-20210317093720-0a41e6d80f39
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.0.0-20210317093720-0a41e6d80f39
	github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3 v3.0.0-20210317093720-0a41e6d80f39
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3 v3.0.0-20210317093720-0a41e6d80f39
	github.com/asim/go-micro/v3 v3.5.0
	github.com/golang/protobuf v1.4.3
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/micro/v3 v3.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	google.golang.org/protobuf v1.25.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
