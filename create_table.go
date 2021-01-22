package main

import (
	"log"
	"os"
)

func createUserInfoTable() {
	err := DB.CreateTable(&UserInfo{}).Error
	if err != nil {
		log.Println("创建数据表失败：", err.Error())
		os.Exit(1)
	}
}
