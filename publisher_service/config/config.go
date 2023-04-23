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
		log.Println("Error while reading config.json")
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		log.Println("Error Unmarrsall", err)
	}
	return config
}

// Client congiguration
func (c Config) ClientConfig() (string, string, string, string) {
	var (
		data      = JsonConfigNodes()
		ws_client = data.Client.WebSocket
		host      = data.Client.Host
		port      = data.Client.Port
		path      = data.Client.Path
	)
	return ws_client, host, port, path
}

// Client congiguration
func (c Config) RolesPermissions() ([]string, []string, []string) {
	var (
		data       = JsonConfigNodes()
		roles      = data.RolesPermissiones.Roles
		permCreate = data.RolesPermissiones.CreateTopic
		permRead   = data.RolesPermissiones.ReadTopic
	)
	return roles, permCreate, permRead
}

// HTTP congiguration
func (c Config) HTTPConfig() string {
	var (
		data = JsonConfigNodes()
		port = data.HTTP.Port
	)
	return port
}

// JSON congiguration
func (c Config) JSONDriversConfig() string {
	var (
		data = JsonConfigNodes()
		path = data.JSONFilesConfig.DriversDemo
	)
	return path
}

// MongoDB configuration
func (c Config) MongoDBConfig() (string, string, string) {
	var (
		data       = JsonConfigNodes()
		dataBase   = data.MongoDBConfig.DataBase
		collection = data.MongoDBConfig.Collection
		mongoUrl   = data.MongoDBConfig.MongoUrl
	)
	return dataBase, collection, mongoUrl
}

// Kafka Config configuration
func (c Config) KafkaConfig() (string, string, string, string) {
	var (
		data              = JsonConfigNodes()
		bootstrap_servers = data.KafkaConfig.BootstrapServers
		clientId          = data.KafkaConfig.ClientId
		acks              = data.KafkaConfig.Acks
		topic             = data.KafkaConfig.Topic
	)
	return bootstrap_servers, clientId, acks, topic
}

// Secret Key Configuration
func (c Config) KeyConfig() string {
	var (
		data = JsonConfigNodes()
		key  = data.SecretKey.SecretKey
	)
	return key
}
