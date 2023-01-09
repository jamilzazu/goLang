package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	db        db
	webserver webserver
	jwt       jwt
}

type db struct {
	Drive    string `mapstructure:"DB_DRIVE"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

type webserver struct {
	Port string `mapstructure:"WEBSERVER_PORT"`
}

type jwt struct {
	Secret    string `mapstructure:"JWT_SECRET"`
	ExpiresIn int    `mapstructure:"JWT_EXPIRE_IN"`
	TokenAuth *jwtauth.JWTAuth
}

func messageError(err error) {
	if err != nil {
		panic(err)
	}
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	messageError(err)

	err = viper.Unmarshal(&cfg)
	messageError(err)

	cfg.jwt.TokenAuth = jwtauth.New("HS256", []byte(cfg.jwt.Secret), nil)
	return cfg, err
}
