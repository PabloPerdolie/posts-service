package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var CONFIG Config

type Config struct {
	Env    string `yaml:"env"`
	Server struct {
		URL string `yaml:"url" env-default:"0.0.0.0:3001"`
	} `yaml:"server"`
	UseInMemory bool `yaml:"useInMemory"`
	DB          struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"db"`
}

func InitConfig() error {
	if err := cleanenv.ReadConfig("config/config.yml", &CONFIG); err != nil {
		log.Fatal(err)
		return err
	}
	err := godotenv.Load()
	if err != nil {
		return err
	}
	if os.Getenv("STORAGE") == "in_memory" {
		CONFIG.UseInMemory = true
	}
	log.Println("initialized config")
	return nil
}
