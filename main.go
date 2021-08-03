package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func dbInit() (*gorm.DB, error) {
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	dsn := fmt.Sprintf("%s:%s%s", user, pass, "@tcp(127.0.0.1:3306)/todolist?charset=utf8&parseTime=True&loc=Local")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {
	act := "run"
	if len(os.Args) == 2 {
		act = os.Args[1]
	}
	db, err := dbInit()
	if err != nil {
		fmt.Println("DB 시작 못함")
	}
	fmt.Println(db)
	if act == "run" {
		fmt.Println("서버 시작")
	} else if act == "migration" {
		fmt.Println("db Init")
	}
}
