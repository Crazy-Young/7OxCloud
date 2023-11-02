package config

type ServiceConfig struct {
	Name string   `json:"name"`
	Host string   `json:"host"`
	Tags []string `json:"tags"`
}
type MySQLConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ServerConfig struct {
	Service ServiceConfig `json:"service"`
	MySQL   MySQLConfig   `json:"mysql"`
	Consul  ConsulConfig  `json:"consul"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataId"`
	Group     string `mapstructure:"group"`
}
