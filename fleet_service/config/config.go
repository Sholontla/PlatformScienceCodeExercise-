package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
}

var Conf = JsonConfigNodes()
var vp *viper.Viper

func JsonConfigNodes() ConfigNodes {
	// Initialize Viper for read config files
	vp = viper.New()
	var config ConfigNodes

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./config_files")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		log.Println("Error while reading init_config.json")
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		log.Println("Error Unmarrsall", err)
	}
	return config
}

// Client congiguration
func (c Config) ServerConfig() (string, string, string, string) {
	var (
		data      = JsonConfigNodes()
		ws_client = data.Server.WebSocket
		host      = data.Server.Host
		port      = data.Server.Port
		path      = data.Server.Path
	)
	return ws_client, host, port, path
}

// HTTP congiguration
func (c Config) HTTPConfig() string {
	var (
		data = JsonConfigNodes()
		port = data.HTTP.Port
	)
	return port
}

// HTTP congiguration
func (c Config) JSONAnalyticsConfig() string {
	var (
		data = JsonConfigNodes()
		path = data.JSONAnalyticsConfig.AnalyticsDemo
	)
	return path
}
