package configs

import "github.com/spf13/viper"

type conf struct {
	IpMaxRequest    int32  `mapstructure:"IP_MAX_REQUESTS"`
	TokenMaxRequest int32  `mapstructure:"TOKEN_MAX_REQUESTS"`
	ApiKey          string `mapstructure:"API_KEY"`
	WebServerPort   string `mapstructure:"WEB_SERVER_PORT"`
	RedisHost       string `mapstructure:"REDIS_HOST"`
	RedisPort       int32  `mapstructure:"REDIS_PORT"`
	RedisPassword   string `mapstructure:"REDIS_PASSWORD"`
	RedisDb         string `mapstructure:"REDIS_DB"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err

}
