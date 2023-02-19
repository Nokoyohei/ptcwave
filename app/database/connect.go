package database

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DBを使い回すことで、DBへのConnectとCloseを毎回しないようにする
var DB *sqlx.DB

// 環境変数を取得する、なければデフォルト値を返す
func getEnv(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

func Connect() error {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.Addr = getEnv("MYSQL_HOST", "localhost") + ":" + getEnv("MYSQL_PORT", "3306")
	config.User = getEnv("MYSQL_USER", "root")
	config.Passwd = getEnv("MYSQL_PASSWORD", "")
	config.DBName = getEnv("MYSQL_DATABASE", "rush_price")
	config.ParseTime = true

	dsn := config.FormatDSN()

	fmt.Println(dsn)

	var err error
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}

	return nil
}
