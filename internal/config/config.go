// internal/config/config.go
package config

// Config root config
type Config struct {
	KIS    KISConfig    `mapstructure:"kis"`
	Server ServerConfig `mapstructure:"server"`
}

type KISConfig struct {
	AppKey       string `mapstructure:"app_key"`
	AppSecret    string `mapstructure:"app_secret"`
	WebsocketURL string `mapstructure:"websocket_url"`
	ApprovalURL  string `mapstructure:"approval_url"`
}

type ServerConfig struct {
	LogLevel string `mapstructure:"log_level"`
}
