package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

type Config struct {
	App      App            `yaml:"app"`
	Server   ServerConfig   `yaml:"server"`
	Log      Log            `yaml:"log"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
}

type App struct {
	Name     string `yaml:"name"`
	Timezone string `yaml:"timezone"`
}

type ServerConfig struct {
	Prefork bool `yaml:"prefork"`
	Port    int  `yaml:"port"`
}

type Log struct {
	Level int `yaml:"level"`
}

type DatabaseConfig struct {
	Username string     `yaml:"username"`
	Password string     `yaml:"password"`
	Host     string     `yaml:"host"`
	Port     int        `yaml:"port"`
	Name     string     `yaml:"name"`
	Pool     PoolConfig `yaml:"pool"`
}

type PoolConfig struct {
	Idle     int `yaml:"idle"`
	Max      int `yaml:"max"`
	Lifetime int `yaml:"lifetime"`
}

type JWTConfig struct {
	Secret     string        `yaml:"secret"`
	Expiration time.Duration `yaml:"expiration"`
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var config Config

	err := v.Unmarshal(&config)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &config, nil
}
