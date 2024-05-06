package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	HttpServer HttpServer `yaml:"http_server"`
	Db         Db         `yaml:"db"`
}

type HttpServer struct {
	Host        string        `yaml:"host" env-default:"localhost"`
	Port        string        `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Db struct {
	User   string `yaml:"user" env-required:"true"`
	Port   int    `yaml:"port" env-required:"true"`
	Name   string `yaml:"name" env-required:"true"`
	Ssl    string `yaml:"ssl" env-default:"disable"`
	Driver string `yaml:"driver" env-required:"true"`
}

func MustLoad() *Config {
	// retrieve config path from env
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	// ensure that config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file '%s' does not exist", configPath)
	}

	// read config file
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Error reading config file '%s': %v", configPath, err)
	}

	return &cfg
}
