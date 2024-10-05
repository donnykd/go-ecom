package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	NAME                   string
	USER                   string
	PASSWORD               string
	HOST                   string
	PORT                   string
	JWTExpirationInSeconds int
	JWTSecret              string
}

func InitConfig() Config {
	exp, err := strconv.Atoi(os.Getenv("JWT_EXP"))
	if err != nil {
		log.Fatal(err)
	}
	return Config{
		NAME:                   os.Getenv("DB_NAME"),
		USER:                   os.Getenv("DB_USER"),
		PASSWORD:               os.Getenv("DB_PASSWORD"),
		HOST:                   os.Getenv("DB_HOST"),
		PORT:                   os.Getenv("DB_PORT"),
		JWTExpirationInSeconds: exp,
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
