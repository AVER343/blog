package config

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Addr     string `json:"address"`
	DbConfig *dbConfig
}

func SetupConfig() (*Config, error) {
	var data = &Config{
		Addr: GetString("ADDR", ":8080"),
		DbConfig: &dbConfig{
			Addr: GetString("DB_ADDR",
				"postgres://postgres:postgres@localhost/postgres?sslmode=disable"),
			MaxIdleConns: GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  GetString("DB_MAX_IDLE_TIME", "15m"),
			MaxOpenConns: GetInt("DB_MAX_OPEN_CONNS", 30),
		},
	}
	return data, nil
}

func PrintConfig(data Config) {
	jsonData, _ := json.Marshal(data)
	fmt.Print(string(jsonData))
}
