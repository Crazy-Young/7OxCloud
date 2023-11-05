package config

type ServiceConfig struct {
	Name string   `json:"name"`
	Host string   `json:"host"`
	Tags []string `json:"tags"`
}

type RedisConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Database   int    `json:"database"`
	Expiration int    `json:"expiration"`
}

type RecommendConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ServerConfig struct {
	Recommend RecommendConfig `json:"recommend"`
	Redis     RedisConfig     `json:"redis"`
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
