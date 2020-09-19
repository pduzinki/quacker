package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Config contains data necessary for app configuration
type Config struct {
	DbConfig       DatabaseConfig `json:"dbConfig"`
	PasswordPepper string         `json:"passwordPepper"`
	HmacKey        string         `json:"hmacKey"`
}

// DatabaseConfig contains data necessary for database configuration
type DatabaseConfig struct {
	Dialect  string `json:"dialect"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// LoadConfig loads and returns Config instance
func LoadConfig() *Config {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println("Failed to read 'config.json'. Loading default config...")
		return LoadDefaultConfig()
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Println("Failed to unmarshal 'config.json'. Loading default config...")
		return LoadDefaultConfig()
	}

	log.Println("Loaded configuration from 'config.json'")
	return &config
}

// GetConnectionInfo returns connectionInfo necessary for establising database connection
func (dbCfg *DatabaseConfig) GetConnectionInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
		"sslmode=disable", dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.Name)
}

// GetDialect returns database dialect necessary for establishing database conection
func (dbCfg *DatabaseConfig) GetDialect() string {
	return dbCfg.Dialect
}

// LoadDefaultConfig returns default config
func LoadDefaultConfig() *Config {
	log.Println("Loaded default configuration")
	return &Config{
		DbConfig: DatabaseConfig{
			Dialect:  "postgres",
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "123",
			Name:     "quacker_dev",
		},
		PasswordPepper: "not-a-secret-pepper",
		HmacKey:        "not-a-secret-hmac-key",
	}
}
