package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jlyon1/IEDFoodStorage/config"
	"github.com/jlyon1/IEDFoodStorage/model"
	"log"
	"time"
)

type Database struct {
	db *sql.DB
}

func (data *Database) Connect(cfg *config.Config) error {
	db, err := sql.Open("mysql",
		cfg.Username+":"+cfg.Password+"@tcp("+cfg.DBHostname+":"+cfg.DBPort+")/food?parseTime=true")
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

	res, err := data.db.Exec("CREATE TABLE pantry(ID int NOT NULL AUTO_INCREMENT KEY, Name varchar(255) NOT NULL, ExpirationDate date, Position int, PadNum int, Count int, Created bigint);")
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
	res, err = data.db.Exec("insert into pantry (ID,Name,Position,PadNum,Count,Created,ExpirationDate) values (1,'Apple',2,1,2,1,01/10/2018);")
}

func (data *Database) GetById(id int) (t model.Food) {
	stmt, err := data.db.Prepare("select * from pantry where `ID` = ?")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&t.ID, &t.Name,
			&t.ExpirationDate, &t.Position,
			&t.PadNum, &t.Count)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
	}
	return t
}

func (data *Database) GetAll() (ret model.Foods) {
	stmt, err := data.db.Prepare("select * from pantry")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		t := model.Food{}
		err := rows.Scan(&t.ID, &t.Name,
			&t.ExpirationDate, &t.Position,
			&t.PadNum, &t.Count, &t.Created)
		if err != nil {
			fmt.Println(err.Error())
		}
		ret.Foods = append(ret.Foods, t)
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
	}
	return ret
}

func (data *Database) Remove(id int) (rem bool) {
	stmt, err := data.db.Prepare("delete from pantry where id=?")
	if err != nil {
		log.Print(err)
		return false
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Print(err)
		return false
	}
	defer rows.Close()
	return true
}

func (data *Database) Update(f model.Food) (ret bool) {
	stmt, err := data.db.Prepare("update pantry set Name=?,Position=?,PadNum=?,Count=?,ExpirationDate=? where ID=? ")
	if f.ID == -1 {
		stmt, err = data.db.Prepare("insert into pantry (Name,Position,PadNum,Count,Created,ExpirationDate) values (?,?,?,?,?,?)")
		_, err = stmt.Exec(f.Name, f.Position, f.PadNum, f.Count, f.Created, time.Now())

	} else {
		_, err = stmt.Exec(f.Name, f.Position, f.PadNum, f.Count, f.ExpirationDate, f.ID)

	}
	if err != nil {
		log.Print(err)
		return false
	}
	defer stmt.Close()

	if err != nil {
		log.Print(err)
		return false
	}
	return true
}
