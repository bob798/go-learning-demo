package main

import (
	"fmt"
	"github.com/bob798/go-learning-demo/go-orm/day1-database-sql/geeorm"
	_ "github.com/mattn/go-sqlite3"
)

//func main() {
//	db, _ := sql.Open("sqlite3", "gee-db")
//	defer func() { _ = db.Close() }()
//	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
//	_, _ = db.Exec("CREATE TABLE  User(Name text);")
//	result, err := db.Exec("INSERT INTO User('Name') values (?),(?)", "Tom", "Sam")
//	if err == nil {
//		affected, _ := result.RowsAffected()
//		log.Println(affected)
//
//	}
//
//	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
//	var name string
//	if err := row.Scan(&name); err == nil {
//		log.Println(name)
//	}
//
//
//}
func main() {
	engine, _ := geeorm.NewEngine("sqlite3", "gee-db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?),(?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected \n ", count)
}
