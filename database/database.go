package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jlyon1/IEDFoodStorage/config"
	"log"
)

type Database struct {
	db *sql.DB
}

func (data *Database) Connect(cfg *config.Config) error {
	db, err := sql.Open("mysql",
		cfg.Username+":"+cfg.Password+"@tcp("+cfg.DBHostname+":"+cfg.DBPort+")/food")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	data.db = db
	return nil
}

func (data *Database) CheckLayout() bool {

	rows, err := data.db.Query("select id from pantry")
	if err != nil {
		fmt.Println("DB config not correct")
		return false
	}

	defer rows.Close()
	err = rows.Err()
	if err != nil {
		fmt.Println("DB config not correct")
		return false
	}
	fmt.Println("DB config correct")

	return true
}

func (data *Database) InitDatabase() {

	res, err := data.db.Exec("CREATE TABLE pantry(ID int NOT NULL AUTO_INCREMENT KEY, Name varchar(255) NOT NULL, ExpirationDate date, Position int, PadNum int);")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}
