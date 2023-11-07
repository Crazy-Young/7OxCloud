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

type RedisConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Database   int    `json:"database"`
	Expiration int    `json:"expiration"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type RabbitMQConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type RecommendConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ServerConfig struct {
	Service   ServiceConfig   `json:"service"`
	MySQL     MySQLConfig     `json:"mysql"`
	Redis     RedisConfig     `json:"redis"`
	Consul    ConsulConfig    `json:"consul"`
	RabbitMQ  RabbitMQConfig  `json:"rabbitmq"`
	Recommend RecommendConfig `json:"recommend"`
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
