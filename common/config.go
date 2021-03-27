package common

import (
	"github.com/asim/go-micro/plugins/config/source/consul/v3"
	"github.com/asim/go-micro/v3/config"
	"strconv"
)

//设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		//设置配置中心地址
		consul.WithAddress(host + ":" + strconv.FormatInt(port, 10)),
		//设置前缀，不设置默认前缀 /micro/config
		consul.WithPrefix(prefix),
		//是否移除前缀，这里设置为true,表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
		)
	//配置初始化
	newConfig, err := config.NewConfig()
	if err != nil {
		return newConfig, err
	}
	err = newConfig.Load(consulSource)
	return newConfig, err
}
