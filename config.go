package main

import "fmt"

// DatabaseConfig contains data necessary for database configuration
type DatabaseConfig struct {
	dialect        string
	connectionInfo string
}

// LoadDatabaseConfig loads and returns DatabaseConfig instance
func LoadDatabaseConfig() *DatabaseConfig {
	// TODO add reading config from file
	return &DatabaseConfig{
		dialect: "postgres",
		connectionInfo: fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
			"sslmode=disable", "localhost", 5432, "postgres", "123", "quacker_dev"),
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
