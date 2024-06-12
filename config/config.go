package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	DBNAME     string
	DBPASSWORD string
	DBUSER     string
	DBPORT     int
	DBHOST     string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return Config{}
	}

	config := Config{}
	config.DBNAME = cast.ToString(get("DBNAME", "secred"))
	config.DBPASSWORD = cast.ToString(get("DBPASSWORD", "secred"))
	config.DBUSER = cast.ToString(get("DBUSER", "secred"))
	config.DBHOST = cast.ToString(get("DBHOST", "google"))
	config.DBPORT = cast.ToInt(get("DBPORT", 9999))

	return config
}

func get(name string, default_value interface{}) interface{} {
	val, exists := os.LookupEnv(name)
	if exists {
		return val
	}
	return default_value
}
