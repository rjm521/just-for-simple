package main

import (
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {

	fmt.Println("hello world")

	// Step 1: 连接数据库
	connect2db()
	// Step 2: 创建数据表
	createUserInfoTable()

	// Step 3: 创建测试数据
	insertRecords()

	// Step 4: 高所你需要的分页类型
	resp, err := getResultSet(1, 5)
	if err != nil {
		log.Println("出错了，您看：", err.Error())
	}

	// Step 5: 把结果输出来 漂亮的json格式有了，还带点颜色，嘿嘿嘿
	s, _ := prettyjson.Marshal(resp)
	fmt.Println(string(s))

	// Step 6: 你想要的啥都有
	fmt.Println(resp.PageNow)    //PageNow: 当前页
	fmt.Println(resp.PageCount)  //PageCount: 总页码数
	fmt.Println(resp.RawCount)   //RawCount: 总行数
	fmt.Println(resp.RawPerPage) //RawPerPage: 每页结果数量
	fmt.Println(resp.ResultSet)  //ResultSet: 返回的数组
	fmt.Println(resp.FirstPage)  //FirstPage: 是否是第一页
	fmt.Println(resp.LastPage)   //LastPage: 是否是最后一页
	fmt.Println(resp.Empty)      //Empty: 该页结果是否为空
	fmt.Println(resp.StartRow)   //StartRow: 本页开始行
	fmt.Println(resp.EndRow)     //EndRow: 本页结束行
}
