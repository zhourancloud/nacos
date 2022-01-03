package main

import (
	"fmt"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

const (
	DataId  = "nacos-simple-demo.yaml"
	GroupId = "DEFAULT_GROUP"
)

func main() {
	// nacos server Config
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "192.168.24.139",
			Port:   8848,
		},
	}

	// 创建客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         "2b1009c4-0d5f-4600-9f5d-1113c1ba0f6d",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./nacos/log",
		CacheDir:            "./nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}

	//获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  GroupId,
	})
	fmt.Println("config: " + content)

	// 监听配置变更
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  GroupId,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group: " + group + ", dataId:" + dataId + ", content" + data)
		},
	})
	if err != nil {
		panic(err)
	}

	for true {
		time.Sleep(1)
	}

}
