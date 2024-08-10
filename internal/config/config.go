package config

type Config struct {
	Engine  *EngineConfig
	Logging *LoggingConfig
	Network *NetworkConfig
}

type EngineConfig struct {
	Type string
}

type LoggingConfig struct {
	Level  string
	Output string
}

type NetworkConfig struct {
	Address        string
	MaxConnections int `mapstructure:"max_connections"`
}
