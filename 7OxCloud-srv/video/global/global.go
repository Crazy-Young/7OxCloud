package global

import (
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/config"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	Debug        bool
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)