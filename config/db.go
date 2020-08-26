package config

type Config struct {
	Host           string
	ConnectionPool int
	DatabaseName   string
}

func New() Config {
	return Config{
		Host:           "mongodb://127.0.0.1:27017",
		ConnectionPool: 5,
		DatabaseName:   "api",
	}
}

func init() {
	New()
}
