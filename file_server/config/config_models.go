package config

type ConfigNodes struct {
	Client          WSClient        `mapstructure:"client"`
	HTTP            HTTPServer      `mapstructure:"http_server"`
	JSONFilesConfig JSONFilesConfig `mapstructure:"drivers_demo"`
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
