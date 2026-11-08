package config

type Config struct {
	Host                string
	Port                int
	LogLevel            string
	InterceptionEnabled bool
}

func DefaultConfig() *Config {
	return &Config{
		Host:                "127.0.0.1",
		Port:                8080,
		LogLevel:            "INFO",
		InterceptionEnabled: false,
	}
}
