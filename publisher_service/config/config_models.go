package config

type ConfigNodes struct {
	RolesPermissiones RolesPermissiones `mapstructure:"users_roles_permissions"`
	Client            WSClient          `mapstructure:"client"`
	HTTP              HTTPServer        `mapstructure:"http_server"`
	JSONFilesConfig   JSONFilesConfig   `mapstructure:"drivers_demo"`
	MongoDBConfig     MongoDBConfig     `mapstructure:"mongo_server"`
	KafkaConfig       KafkaConfig       `mapstructure:"kafka_config"`
	SecretKey         SecretKey         `mapstructure:"secret_key"`
}

type WSClient struct {
	WebSocket string `mapstructure:"ws_client"`
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	Path      string `mapstructure:"path"`
}

type RolesPermissiones struct {
	Roles       []string `mapstructure:"roles"`
	CreateTopic []string `mapstructure:"create_topic"`
	ReadTopic   []string `mapstructure:"read_topic"`
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
	MongoUrl   string `mapstructure:"mongo_url"`
}

type KafkaConfig struct {
	BootstrapServers string `mapstructure:"bootstrap_servers"`
	ClientId         string `mapstructure:"client_id"`
	Acks             string `mapstructure:"acks"`
	Topic            string `mapstructure:"topic"`
}

type SecretKey struct {
	SecretKey string `mapstructure:"key"`
}
