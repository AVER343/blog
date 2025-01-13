package config

type dbConfig struct {
	Addr         string
	MaxIdleConns int
	MaxOpenConns int
	MaxIdleTime  string
}
