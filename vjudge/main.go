package vjudge

import (
	"encoding/json"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/imroc/req"
	"github.com/intxiaoquan/vjudge_lab_helper/handle"
	"github.com/intxiaoquan/vjudge_lab_helper/jsonstruct"
	"github.com/intxiaoquan/vjudge_lab_helper/util"
)

var (
	problemNum      int
	contestData     jsonstruct.ContestInfo
	descriptionData jsonstruct.DescriptionInfo
	codeQueryData   jsonstruct.CodeQueryInfo
	problemData     [20]jsonstruct.ProblemInfo //储存每个题目各个字段的结构体
	outContent      string                     //临时存储单个实验问题描述
	outCode         string                     //临时存储单个实验AC代码
)

const (
	loginUrl       = "https://vjudge.net/user/login"
	contestUrl     = "https://vjudge.net/contest/"
	descriptionUrl = "https://vjudge.net/problem/description/"
	reqCodeUrl     = "https://vjudge.net/solution/data/"
)

func Login(username string, password string) (cookie []*http.Cookie, err error) {
	log.Println("正在登陆...")
	header := req.Header{
		"Accept":       "*/*",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		"Host":         "vjudge.net",
	}
	param := req.Param{
		"username": username,
		"password": password,
	}

	r, err := req.Post(loginUrl, header, param)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("%+v", r.Response().Body)
	cookie = r.Response().Cookies()
	return cookie, nil
}

func GetData(username string, cookie []*http.Cookie, contestID string, contestNum int, ans *[20]jsonstruct.Output2File) {

	log.Println("开始爬取实验数据...")
	r, err := req.Get(contestUrl+contestID, cookie)
	if err != nil {
		log.Fatal(err)
	}
	up := "<textarea style=\"display: none\" name=\"dataJson\">"
	down := "</textarea>"
	jsonData := util.Between(r.String(), up, down)
	json.Unmarshal([]byte(jsonData), &contestData)
	problemNum = len(contestData.Problems)
	log.Println("[实验查询成功] 实验名:" + contestData.Title + " 题目个数:" + strconv.Itoa(problemNum))

	log.Println("开始遍历题目...")
	//遍历题目数查询题目详细信息
	for i, item := range contestData.Problems {
		log.Println("[正在处理题目] " + problemData[i].Tag + "-" + problemData[i].Title)
		id := item.PublicDescID
		t := item.PublicDescVersion
		r, err = req.Get(descriptionUrl+strconv.Itoa(id)+"?"+strconv.FormatInt(t, 10), cookie)
		if err != nil {
			log.Fatal(err)
		}
		up := "<textarea class=\"data-json-container\" style=\"display: none\">"
		down := "</textarea>"
		jsonData := util.Between(r.String(), up, down)
		content := strings.Replace(jsonData, "\\u003c", "<", -1)
		content = strings.Replace(content, "\\u003e", ">", -1)
		content = strings.Replace(content, "\\u0026", "&", -1)
		content = util.TrimHtml(content)
		flag := strings.Contains(content, "\\r\\n")
		if flag {
			content = strings.Replace(content, "\\n      ", "\\r\\n", -1)
		} else {
			content = strings.Replace(content, "\\n   ", "", -1)
			content = strings.Replace(content, "\\n", "\\r\\n", -1)
		}
		jsonData = content
		json.Unmarshal([]byte(jsonData), &descriptionData)
		//开始生成题目信息结构体
		var targetChar rune
		targetChar = rune(65 + i)
		problemData[i].Tag = string(targetChar)
		problemData[i].Title = item.Title
		problemData[i].Description = html.UnescapeString(descriptionData.Sections[0].Value.Content)
		problemData[i].Input = html.UnescapeString(descriptionData.Sections[1].Value.Content)
		problemData[i].Output = html.UnescapeString(descriptionData.Sections[2].Value.Content)
		problemData[i].SampleInput = html.UnescapeString(descriptionData.Sections[3].Value.Content)
		problemData[i].SampleOutput = html.UnescapeString(descriptionData.Sections[4].Value.Content)
		log.Println("[题目信息查询成功] " + problemData[i].Tag + "-" + problemData[i].Title)
		//获取AC runID
		submitUrl := "https://vjudge.net/status/data/?draw=1&start=0&length=1&un=" + username + "&num=" + string(targetChar) + "&res=1&language=&inContest=true&contestId=" + contestID

		log.Println("[开始获取AC代码状态码]" + problemData[i].Tag + "-" + problemData[i].Title)
		r, err = req.Get(submitUrl, cookie)
		if err != nil {
			log.Fatal(err)
		}
		up = "{\"data\":["
		down = "],\"recordsTotal\""
		jsonData = util.Between(r.String(), up, down)
		json.Unmarshal([]byte(jsonData), &codeQueryData)

		log.Println("[AC状态码获取成功] 状态码:" + strconv.Itoa(codeQueryData.RunID))
		//获取AC代码并拼接题目结构体
		log.Println("[开始查询AC代码]:" + problemData[i].Tag + "-" + problemData[i].Title)
		r, err = req.Get(reqCodeUrl+strconv.Itoa(codeQueryData.RunID), cookie)
		if err != nil {
			log.Fatal(err)
		}
		var reqCodeData jsonstruct.ReqCodeInfo
		jsonData = r.String()
		content = strings.Replace(jsonData, "\\u003c", "<", -1)
		content = strings.Replace(content, "\\u003e", ">", -1)
		content = strings.Replace(content, "\\u0026", "&", -1)
		content = strings.Replace(content, "\\n", "\\r\\n", -1)
		jsonData = content
		json.Unmarshal([]byte(jsonData), &reqCodeData)
		problemData[i].Code = reqCodeData.Code
		log.Println("[AC代码查询成功] 题目:" + problemData[i].Tag + "-" + problemData[i].Title)
		//写代码到文件
		log.Println("[开始生成代码文件] 题目:" + problemData[i].Tag + "-" + problemData[i].Title)
		handle.FileOn(username, contestNum, problemNum, problemData[i].Code)
		log.Println("[生成代码文件结束] 题目:" + problemData[i].Tag + "-" + problemData[i].Title)
		problemNum = i + 1
		outContent += "Title:" + problemData[i].Tag + "-" + problemData[i].Title
		outContent += "\r\nDescription:\r\n" + problemData[i].Description
		outContent += "\r\nInput:\r\n" + problemData[i].Input
		outContent += "\r\nOutput:\r\n" + problemData[i].Output
		outContent += "\r\nSampleInput:\r\n " + problemData[i].SampleInput
		outContent += "\r\nSampleOutput:\r\n" + problemData[i].SampleOutput + "\r\n\r\n"
		outCode += "Title: " + problemData[i].Tag + "-" + problemData[i].Title
		outCode += "\r\nCode: \r\n" + problemData[i].Code + "\r\n\r\n"

		//结果写入数组
		log.Println("[数据写入数组] 题目:" + problemData[i].Tag + "-" + problemData[i].Title)
		(*ans)[contestNum].Content = outContent
		(*ans)[contestNum].Code = outCode
	}
	log.Println("[实验数据处理完成] 实验名:" + contestData.Title + " 题目个数:" + strconv.Itoa(problemNum))
}
