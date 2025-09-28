package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Adress string `yaml:"address" env-required:"true"`
}

// We'll give structs tags to the entities of struct to extract contents from yaml file

type Config struct {
	Env string `yaml:"env" env:"dev" env-required:"true"`

	StoragePath string `yaml:"storage_path" env:"storage.db" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "Path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config path is not set")
		}
	}
	// setting the config path checker if file doesn't exist the program will stop right away trying it with file handling concept

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("failed to read config file: %s", err.Error())
	}
	return &cfg
}
