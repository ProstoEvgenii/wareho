package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	PG     PostgresConfig
	Server Server
}

type PostgresConfig struct {
	DSN string
}

type Server struct {
	ServerAddr string
}

func ParseEnv() *Config {
	conf := &Config{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Err(err).Send()
	}

	// postgreSQL config
	conf.PG = PostgresConfig{
		DSN: viper.GetString("PG_DSN"),
	}

	// fasthttpServer config
	conf.Server = Server{
		//Debug:      viper.GetBool("DEBUG"),
		ServerAddr: viper.GetString("SERVER_ADDR"),
	}

	return conf
}
