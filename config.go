package main

import "fmt"

type DatabaseConfig struct {
	dialect        string
	connectionInfo string
}

func LoadDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		dialect: "postgres",
		connectionInfo: fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
			"sslmode=disable", "localhost", 5432, "postgres", "123", "quacker_dev"),
	}
}

func (dbCfg *DatabaseConfig) ConnectionInfo() string {
	return dbCfg.connectionInfo
}

func (dbCfg *DatabaseConfig) Dialect() string {
	return dbCfg.dialect
}
