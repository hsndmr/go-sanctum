package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database *Database
	EnvType string
	Port string
	Debug bool
}


// Init return the config
func Init() (*Config, error) {
	env, err := FindEnv()
	if err != nil {
		return nil, err
	}

	err = godotenv.Load(env)
	if err != nil {
    return nil, err
  }

	config := &Config{}

	config.Database = databaseSetting()

	config.EnvType = os.Getenv("APP_ENV")

	config.Port = os.Getenv("APP_PORT")

	config.Debug = os.Getenv("APP_DEBUG") == "true"

	return config, nil
}


// Database is the database config
type Database struct {
	Connection string
	Host     string
	Name     string
	User     string
	Password string
}


// databaseSetting returns the database config
func databaseSetting() *Database {
	databaseConfig := &Database{}
	databaseConfig.Connection = os.Getenv("DB_CONNECTION")
	databaseConfig.Host = os.Getenv("DB_HOST")
	databaseConfig.Name = os.Getenv("DB_DATABASE")
	databaseConfig.User = os.Getenv("DB_USERNAME")
	databaseConfig.Password = os.Getenv("DB_PASSWORD")	
	return databaseConfig
}

// isEnvExists checks if the file exists
func isExistFile(file string) bool {
	_, err := os.Stat(file);
	return err == nil
}

// isRunningOnTest checks if the application is running on test
func isRunningOnTest() bool {
	return flag.Lookup("test.v") != nil
}
 
// findEnvFile finds the env file
func FindEnv() (string, error) {
	file := ".env"

	if isRunningOnTest() {
		file += ".testing"
	}

	if(isExistFile("./"+file)) {
		return "./"+file, nil
	}

	dept := "../"

	for i := 1; i <= 10; i++ {
		if(isExistFile(dept+file)) {
			return dept+file, nil
		}
		dept += "../"
	}
	
	return "", fmt.Errorf("env file not found")
}