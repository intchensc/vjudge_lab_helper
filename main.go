package main

import (
	"fmt"
	"log"

	Vjudge "github.com/intxiaoquan/vjudge_lab_helper/vjudge"
)

const (
	problemMaxNum = 20
)

func main() {
	var username string
	var password string
	var contestID string

	username = "071219130"
	password = "12175210csc"
	contestID = "442736"

	fmt.Println("请输入vjudge的账号和密码，用空格隔开！")
	// fmt.Scanln(&username, &password)
	//登陆获取cookie
	cookie, err := Vjudge.Login(username, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("请输入实验编号：")

	//获取实验信息
	Vjudge.GetData(username, cookie, contestID)
}
