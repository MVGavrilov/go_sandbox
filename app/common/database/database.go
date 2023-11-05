package database

import (
	"github.com/jmoiron/sqlx"
	"strconv"
	"log"

	_ "github.com/lib/pq"
)

var (
  Handler *sqlx.DB
)

type PostgresConfig struct {
	Host      string  `default:"127.0.0.1"`
  Port      int     `default:"5432"`
  Database  string  `default:""`
	User      string  `default:"user"`
  Password  string  `default:"password"`
}

func ParseConfig(config PostgresConfig) string {
  ret := "host=" + config.Host + " " +
          "port=" + strconv.Itoa(config.Port) + " " +
          "user=" + config.User + " " +
          "password=" + config.Password + " "
  if config.Database != "" {
    ret += "dbname=" + config.Database + " "
  }
  ret += "sslmode=disable"
  return ret
}

func Connect(config PostgresConfig) error {
  var err error

  if Handler, err = sqlx.Open("postgres", ParseConfig(config)); err != nil {
    log.Println("SQL Driver error", err)
  }

  if err = Handler.Ping(); err != nil {
    log.Println("Database error", err)
  }

  return err
}

func Close() {
  Handler.Close()
}