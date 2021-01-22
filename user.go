package main

import "github.com/jinzhu/gorm"

type UserInfo struct {
	gorm.Model
	UserName string
	Age      int
	Gender   string
	School   string
	Home     string
}
