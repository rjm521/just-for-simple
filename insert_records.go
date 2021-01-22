package main

import (
	"fmt"
	"log"
)

func insertRecords() {
	user := UserInfo{}
	for i := 0; i < 100; i++ {
		user.ID = uint(i)
		user.UserName = fmt.Sprintf("test)->%d", i+1)
		if i%2 == 0 {
			user.Age = 1
		} else {
			user.Age = 2
		}
		user.Gender = "m"
		user.Home = "weinan"
		user.School = "njupt-iot"
		if err := DB.Model(&UserInfo{}).Create(&user).Error; err != nil {
			log.Println("创建用户", i, "失败:", err.Error())
		}
	}
	log.Println("您想要的数据已经生成成功")
}
