package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func connect2db() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "tech_blog_system")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("connect database failed: ", err.Error())
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n[DB LOG] ", log.Lshortfile))
	DB = db
}
