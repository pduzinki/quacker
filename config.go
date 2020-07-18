package main

import "fmt"

// Config contains data necessary for app configuration
type Config struct {
	DbConfig       DatabaseConfig
	PasswordPepper string
}

// DatabaseConfig contains data necessary for database configuration
type DatabaseConfig struct {
	dialect        string
	connectionInfo string
}

// LoadConfig loads and returns Config instance
func LoadConfig() *Config {
	// TODO add reading config from file
	return &Config{
		DbConfig: DatabaseConfig{
			dialect: "postgres",
			connectionInfo: fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
				"sslmode=disable", "localhost", 5432, "postgres", "123", "quacker_dev"),
		},
		PasswordPepper: "not-a-secret-pepper",
	}
}

// ConnectionInfo returns connectionInfo necessary for establising database connection
func (dbCfg *DatabaseConfig) ConnectionInfo() string {
	return dbCfg.connectionInfo
}

// Dialect returns database dialect necessary for establishing database conection
func (dbCfg *DatabaseConfig) Dialect() string {
	return dbCfg.dialect
}
