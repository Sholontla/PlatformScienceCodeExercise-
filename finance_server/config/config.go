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

func GRPCConfig() (string, string, bool) {
	data := JsonConfigNodes()
	grpc_host := data.GRPCServer.GRPCHost
	grpc_port := data.GRPCServer.GRPCPort
	grpc_tls := data.GRPCServer.GRPCTLS

	return grpc_host, grpc_port, grpc_tls
}

func GRPCPathsConfig() (string, string) {
	data := JsonConfigNodes()
	cert_File := data.GRPCPaths.CertFile
	key_File := data.GRPCPaths.CertFile

	return cert_File, key_File
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
		driver     = data.MongoDBConfig.Driver
	)
	return collection, dataBase, driver
}
