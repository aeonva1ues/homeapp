package infrastructure

import (
	"github.com/spf13/viper"
)

type Config interface {
	GetServerHost() string
	GetDatabase() *Database
	GetUploadsDir() string
}

type Server struct {
	Host string
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Uploads struct {
	Path string
}

type YamlConfig struct {
	Server   Server
	Database Database
	Uploads  Uploads
}

func (y *YamlConfig) GetServerHost() string {
	return y.Server.Host
}

func (y *YamlConfig) GetDatabase() *Database {
	return &y.Database
}

func (y *YamlConfig) GetUploadsDir() string {
	return y.Uploads.Path
}

type EnvConfig struct {
	Host             string `mapstructure:"HOST"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     int    `mapstructure:"DATABASE_PORT"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	SSLMode          string `mapstructure:"SSL_MODE"`
	Uploads          string `mapstructure:"UPLOADS"`
}

func (e *EnvConfig) GetServerHost() string {
	return e.Host
}

func (e *EnvConfig) GetDatabase() *Database {
	return &Database{
		Host:     e.DatabaseHost,
		Port:     e.DatabasePort,
		User:     e.DatabaseUser,
		Password: e.DatabasePassword,
		DBName:   e.DatabaseName,
		SSLMode:  e.SSLMode,
	}
}

func (e *EnvConfig) GetUploadsDir() string {
	return e.Uploads
}

func LoadEnvConfig() (Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg EnvConfig
	viper.Unmarshal(&cfg)
	return &cfg, nil
}

func LoadYamlConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg YamlConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
