package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
  "github.com/jlyon1/IEDFoodStorage/config"
  "log"
)

type Database struct{
  db mysql.MySQLDriver
}

func (data *Database) Connect(cfg *config.Config) error{
	db, err := sql.Open("mysql",
		cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.DBHostname + ":" + cfg.DBPort + ")/food")
	if err != nil {
		log.Fatal(err)
    return err
	}
	defer db.Close()
  return nil
}
