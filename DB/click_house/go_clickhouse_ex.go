package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func mainn() {
	conn, err := connect()
	if err != nil {
		panic((err))
	}

	ctx := context.Background()
	rows, err := conn.Query(ctx, "SELECT name,toString(uuid) as uuid_str FROM system.tables LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			name, uuid string
		)
		if err := rows.Scan(
			&name,
			&uuid,
		); err != nil {
			log.Fatal(err)
		}
		log.Printf("name: %s, uuid: %s",
			name, uuid)
	}

}

func connect() (driver.Conn, error) {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{

			Addr: []string{"lqlzygbi5p.europe-west4.gcp.clickhouse.cloud:8443"},

			Auth: clickhouse.Auth{
				Database: "default",
				Username: "default",
				Password: "wCf~Al.XbECM2",
			},

			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "an-example-go-client", Version: "0.1"},
				},
			},

			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
			TLS: &tls.Config{
				InsecureSkipVerify: true,
			},
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}

	fmt.Println("connected")
	return conn, nil
}

func Conn() {
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"lqlzygbi5p.europe-west4.gcp.clickhouse.cloud:8443"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "wCf~Al.XbECM2",
		},
		TLS: &tls.Config{InsecureSkipVerify: true},
	})

	fmt.Println("connection:", conn)

	row := conn.QueryRow(`SELECT username FROM "users" LIMIT 1`)
	var col string
	if err := row.Scan(&col); err != nil {
		fmt.Printf("An error while reading the data: %s", err)
	} else {
		fmt.Printf("Result: %s", col)
	}
}

func Conn1() {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"lqlzygbi5p.europe-west4.gcp.clickhouse.cloud:8443"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "wCf~Al.XbECM2",
		},
		TLS: &tls.Config{},
	})
	if err != nil {
		println("1", err)
		return
	}
	ctx := context.Background()

	rows, err := conn.Query(ctx, "SELECT username FROM users")
	if err != nil {
		println("2", err.Error())
		return
	}
	for rows.Next() {
		var (
			col2 string
		)
		if err := rows.Scan(&col2); err != nil {
			println("3", err)
			return
		}
		fmt.Printf("row: col2=%s\n", col2)
	}

	defer rows.Close()

	fmt.Println(conn)

}
func main() {
	Conn1()
}
