package initialize

import (
	"encoding/json"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/global"
	"github.com/spf13/viper"
)

func InitConfig() {
	configFileName := "7OxCloud-api/config.yaml"
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic("读取配置文件失败：" + err.Error())
	}
	if err := v.Unmarshal(&global.NacosConfig); err != nil {
		panic("配置反序列化失败：" + err.Error())
	}

	global.Debug = false
	logLevel := "warn"
	if global.NacosConfig.Group == "dev" {
		global.Debug = true
		logLevel = "info"
	}
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   uint64(global.NacosConfig.Port),
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "7OxCloud-api/tmp/nacos/log",
		CacheDir:            "7OxCloud-api/tmp/nacos/cache",
		LogLevel:            logLevel,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic("创建nacos客户端失败：" + err.Error())
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})
	if err != nil {
		panic("获取nacos配置失败：" + err.Error())
	}
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		panic("配置反序列化失败：" + err.Error())
	}
}
