package config

type ApiConfig struct {
	Name string   `json:"name"`
	Host string   `json:"host"`
	Tags []string `json:"tags"`
}
type ServiceConfig struct {
	User  string `json:"user"`
	Video string `json:"video"`
}

type JWTConfig struct {
	SigningKey string `json:"signingKey"`
	Expiration int    `json:"expiration"`
}

type AliSmsConfig struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	SignName        string `json:"signName"`
	TemplateCode    string `json:"templateCode"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type RedisConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Expiration int    `json:"expiration"`
}

type QiNiuYunConfig struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
	Domain    string `json:"domain"`
}

type ServerConfig struct {
	Api      ApiConfig      `json:"api"`
	Service  ServiceConfig  `json:"service"`
	JWT      JWTConfig      `json:"jwt"`
	AliSms   AliSmsConfig   `json:"sms"`
	Redis    RedisConfig    `json:"redis"`
	Consul   ConsulConfig   `json:"consul"`
	QiNiuYun QiNiuYunConfig `mapstructure:"qiniuyun"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"userService"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataId"`
	Group     string `mapstructure:"group"`
}
