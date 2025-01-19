package config

import (
	"strconv"
)

var env map[string]string = make(map[string]string, 0)

func init() {
	env["ADDR"] = GetString("ADDR", ":8080")
	env["DB_ADDR"] = GetString("DB_ADDR", "postgresql://AVER343:NCFH4lqyM2nz@ep-small-salad-03265560.ap-southeast-1.aws.neon.tech/netflix_database?sslmode=require")
	env["DB_MAX_IDLE_CONNS"] = strconv.Itoa(GetInt("DB_MAX_IDLE_CONNS", 30))
	env["DB_MAX_IDLE_TIME"] = GetString("DB_MAX_IDLE_TIME", "15m")
	env["DB_MAX_OPEN_CONNS"] = strconv.Itoa(GetInt("DB_MAX_OPEN_CONNS", 30))
}

func GetString(key, fallback string) string {
	val, ok := env[key]
	if !ok {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val, ok := env[key]
	if !ok {
		return fallback
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return intVal
}
