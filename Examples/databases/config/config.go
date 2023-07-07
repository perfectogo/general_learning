package config

import _ "github.com/ClickHouse/clickhouse-go"

func LoadConfig() *Config {

	var config = &Config{
		ClickHouseConfig: ClickHouseConfig{
			Host:     "localhost",
			Port:     9000,
			Username: "default",
			Password: "",
			Database: "default",
		},
	}

	return config
}
