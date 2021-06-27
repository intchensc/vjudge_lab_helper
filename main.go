package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/howeyc/gopass"
	"github.com/intxiaoquan/vjudge_lab_helper/handle"
	"github.com/intxiaoquan/vjudge_lab_helper/jsonstruct"
	"github.com/intxiaoquan/vjudge_lab_helper/vjudge"
)

func main() {
	var username string                    //用户名
	var password string                    //密码
	var contestID [20]string               //实验ID数组，默认最大值为20
	var outData [20]jsonstruct.Output2File //写入word结果的结构体
	var cntContest int
	fmt.Println()
	fmt.Println("-----Vjudge-Lab-Helper V0.9-----")
	fmt.Println("--------- @intxiaoquan----------")
	fmt.Println()

	fmt.Println("请输入vjudge的账号:")
	fmt.Scanln(&username)
	fmt.Println("请输入vjudge的密码(命令行不显示):")
	pass, err := gopass.GetPasswd()
	password = string(pass)
	//登陆获取cookie
	cookie, err := vjudge.Login(username, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("请输入实验个数:")
	fmt.Scanln(&cntContest)
	fmt.Println("请输入实验ID, 使用空格隔开！:")
	for i := 0; i < cntContest; i++ {
		fmt.Scanf("%s", &contestID[i])
	}

	//获取实验信息
	for i := 0; i < cntContest; i++ {
		var p *[20]jsonstruct.Output2File = &outData
		vjudge.GetData(username, cookie, contestID[i], i, p)

	}
	log.Println("[开始写入word]")
	handle.DocOn(outData, cntContest)
	log.Println("[写入word结束]")
	log.Println("[程序运行结束] 共完成实验:" + strconv.Itoa(cntContest) + "个\n")
}
