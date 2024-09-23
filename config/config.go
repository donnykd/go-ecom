package config

import "os"

type Config struct {
	NAME     string
	USER     string
	PASSWORD string
	HOST     string
	PORT     string
}

func InitConfig() Config{
	return Config{
		NAME:     os.Getenv("DB_NAME"),
		USER:     os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		HOST:     os.Getenv("DB_HOST"),
		PORT:     os.Getenv("DB_PORT"),
	}
}
