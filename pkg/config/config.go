package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EnvDepth int
	Database Database
	Env string
}

// InitEnv initializes the env file
func (c *Config) initEnv() {
	env, err := findEnv()
	if err != nil {
		log.Fatal("Not found .env file")
	}

	err = godotenv.Load(env)
	if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func (c *Config) initApp() {
	c.initEnv()
	c.Env = os.Getenv("APP_ENV")
}

// Init initializes the config
func (c *Config) Init() {
	c.initApp()
	c.initDatabaseSetting()
}

// App is the global config
var App = &Config{
	EnvDepth: 20,
}

// Database is the database config
type Database struct {
	Connection string
	Host     string
	Name     string
	User     string
	Password string
}

// DatabaseConfig is the database config
var DatabaseConfig = &Database{}

// initDatabaseSetting initializes the database setting
func (c *Config) initDatabaseSetting() {
	DatabaseConfig.Connection = os.Getenv("DB_CONNECTION")
	DatabaseConfig.Host = os.Getenv("DB_HOST")
	DatabaseConfig.Name = os.Getenv("DB_DATABASE")
	DatabaseConfig.User = os.Getenv("DB_USERNAME")
	DatabaseConfig.Password = os.Getenv("DB_PASSWORD")	
	App.Database = *DatabaseConfig
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
func findEnv() (string, error) {
	file := ".env"

	if isRunningOnTest() {
		file += ".testing"
	}

	if(isExistFile("./"+file)) {
		return "./"+file, nil
	}

	dept := "../"

	for i := 1; i <= App.EnvDepth; i++ {
		if(isExistFile(dept+file)) {
			return dept+file, nil
		}
		dept += "../"
	}
	
	return "", fmt.Errorf("env file not found")
}