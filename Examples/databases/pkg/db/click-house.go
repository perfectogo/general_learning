package db

import (
	"database/sql"
	"examples/databases/config"
	"fmt"
)

func GetClickHouseConnection(config *config.Config) (*sql.DB, error) {

	connectionString := fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&database=%s",
		config.ClickHouseConfig.Host,
		config.ClickHouseConfig.Port,
		config.ClickHouseConfig.Username,
		config.ClickHouseConfig.Password,
		config.ClickHouseConfig.Database,
	)

	sqlDB, err := sql.Open("clickhouse", connectionString)
	if err != nil {
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return nil, err
	}

	return sqlDB, nil
}
