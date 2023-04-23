package config

type ConfigNodes struct {
	Client          WSClient        `mapstructure:"client"`
	HTTP            HTTPServer      `mapstructure:"http_server"`
	JSONFilesConfig JSONFilesConfig `mapstructure:"drivers_demo"`
	MongoDBConfig   MongoDBConfig   `mapstructure:"mongo_server"`
	GRPCServer      GRPCServer      `mapstructure:"grpc_server"`
	GRPCPaths       GRPCPaths       `mapstructure:"config_grpc_path"`
}

type GRPCPaths struct {
	CertFile string `mapstructure:"cert_File"`
	KeyFile  string `mapstructure:"key_File"`
}

type GRPCServer struct {
	GRPCHost string `mapstructure:"host"`
	GRPCPort string `mapstructure:"port"`
	GRPCTLS  bool   `mapstructure:"tls"`
}

type WSClient struct {
	WebSocket string `mapstructure:"ws_client"`
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	Path      string `mapstructure:"path"`
}

type HTTPServer struct {
	Port string `mapstructure:"port"`
}

type JSONFilesConfig struct {
	DriversDemo string `mapstructure:"drivers_demo"`
}

type MongoDBConfig struct {
	DataBase   string `mapstructure:"data_base"`
	Collection string `mapstructure:"collection"`
	Driver     string `mapstructure:"driver"`
}
