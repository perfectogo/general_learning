package config

type Config struct {
	PostgresConfig   PostgresConfig
	ClickHouseConfig ClickHouseConfig
	MongoConfig      MongoConfig
	RedisConfig      RedisConfig
}

type PostgresConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type ClickHouseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type MongoConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}
