package config

type ConfigNodes struct {
	Server              WSServer            `mapstructure:"server"`
	HTTP                HTTPServer          `mapstructure:"http_server"`
	JSONAnalyticsConfig JSONAnalyticsConfig `mapstructure:"analytic_server"`
}

type WSServer struct {
	WebSocket string `mapstructure:"ws_server"`
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	Path      string `mapstructure:"path"`
}

type HTTPServer struct {
	Port string `mapstructure:"port"`
}

type JSONAnalyticsConfig struct {
	AnalyticsDemo string `mapstructure:"analytic_server"`
}
