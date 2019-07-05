package postgres

import (
	"../../types"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

var DBX *sqlx.DB

func InitX(myConfig types.Env) {
	var err error
	var dataSource = "postgres://" + myConfig.PostgresUser + ":" + myConfig.PostgresPassword + "@" + myConfig.PostgresHost + ":" + myConfig.PostgresPort + "/" + myConfig.PostgresBase + "?sslmode=disable"
	DBX, err = sqlx.Connect("postgres", dataSource)
	if err == nil {
		fmt.Println("\x1b[32mPostgres connect OK\x1b[0m")
	} else {
		fmt.Println("\x1b[31mPostgres connect ERROR\x1b[0m", err)
		os.Exit(10)
	}
	DBX.SetConnMaxLifetime(0)
	DBX.SetMaxIdleConns(5)
	DBX.SetMaxOpenConns(10)
	if err = DBX.Ping(); err != nil {
		os.Exit(10)
	}
}

func Check() bool {
	err := DBX.Ping()
	if err != nil {
		return false
	}
	return true
}
