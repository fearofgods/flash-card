package setting

type Config struct {
	Logger  LoggerSetting `mapstructure:"logger"`
	AppInfo AppInfo       `mapstructure:"app"`
	Server  ServerSetting `mapstructure:"server"`
}

type LoggerSetting struct {
	LogLevel   string `mapstructure:"level"`
	FileFormat string `mapstructure:"file_format"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type ServerSetting struct {
	Port string `mapstructure:"port"`
}

type AppInfo struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Env     string `mapstructure:"env"`
}
