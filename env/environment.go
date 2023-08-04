package env

import (
	"github.com/spf13/viper"
	"strings"
)

var Env = &Environment{}

type Environment struct {
	App struct {
		Name string `mapstructure:"NAME"`
		Port int    `mapstructure:"PORT"`
	} `mapstructure:"APP"`
	JWT struct {
		JwtSecret                string `mapstructure:"SECRET"`
		JwtIssuer                string `mapstructure:"ISSUER"`
		JwtDomain                string `mapstructure:"DOMAIN"`
		JwtExpirationTime        int64  `mapstructure:"EXPIRATION_TIME"`
		JwtRefreshExpirationTime int64  `mapstructure:"REFRESH_EXPIRATION_TIME"`
	} `mapstructure:"JWT"`
	Database struct {
		DSN string `mapstructure:"DSN"`
	} `mapstructure:"DATABASE"`
}

func ReadConfig(path string) (*Environment, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(path)
	v.SetConfigType("yml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&Env); err != nil {
		return nil, err
	}

	return Env, nil
}
