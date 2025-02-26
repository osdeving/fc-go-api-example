package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)



type conf struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	JwtSecret string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn int `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	
	viper.SetConfigFile(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JwtSecret), nil)
	return cfg, err
}
