package Vjudge

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req"
	"github.com/intxiaoquan/vjudge_lab_helper/Handledocx"
	Jsonstruct "github.com/intxiaoquan/vjudge_lab_helper/JsonStruct"
	"github.com/intxiaoquan/vjudge_lab_helper/Util"
	Handlefile "github.com/intxiaoquan/vjudge_lab_helper/handleFile"
)

var (
	problemNum      int
	contestData     Jsonstruct.ContestInfo
	descriptionData Jsonstruct.DescriptionInfo
	codeQueryData   Jsonstruct.CodeQueryInfo
	problemData     [20]Jsonstruct.ProblemInfo
	outContent      string
	outCode         string
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
	log.Printf("%+v", r)
	cookie = r.Response().Cookies()
	return cookie, nil
}

func GetData(username string, cookie []*http.Cookie, contestID string) {
	r, err := req.Get(contestUrl+contestID, cookie)
	if err != nil {
		log.Fatal(err)
	}
	up := "<textarea style=\"display: none\" name=\"dataJson\">"
	down := "</textarea>"
	jsonData := Util.Between(r.String(), up, down)
	json.Unmarshal([]byte(jsonData), &contestData)

	problemNum = len(contestData.Problems)
	fmt.Printf("len:%d\n", problemNum)
	//获取文字描述
	fmt.Printf("实验名称：%s\n", contestData.Title)
	//遍历题目数组获取id和时间戳用于查询题目详细信息
	for i, item := range contestData.Problems {
		id := item.PublicDescID
		t := item.PublicDescVersion
		r, err = req.Get(descriptionUrl+strconv.Itoa(id)+"?"+strconv.FormatInt(t, 10), cookie)
		if err != nil {
			log.Fatal(err)
		}

		up := "<textarea class=\"data-json-container\" style=\"display: none\">"
		down := "</textarea>"
		jsonData := Util.Between(r.String(), up, down)

		content := strings.Replace(jsonData, "\\u003c", "<", -1)
		content = strings.Replace(content, "\\u003e", ">", -1)
		content = strings.Replace(content, "\\u0026", "&", -1)
		content = Util.TrimHtml(content)
		flag := strings.Contains(content, "\\r\\n")
		if flag {
			content = strings.Replace(content, "\\n      ", "\\r\\n", -1)
		} else {
			content = strings.Replace(content, "\\n   ", "", -1)
			content = strings.Replace(content, "\\n", "\\r\\n", -1)
		}
		jsonData = content
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(r.String()))
		fmt.Println(r.String())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dom)
		dom.Find("dd").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Text())
		})

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
		//获取AC runID
		submitUrl := "https://vjudge.net/status/data/?draw=1&start=0&length=1&un=" + username + "&num=" + string(targetChar) + "&res=1&language=&inContest=true&contestId=" + contestID
		fmt.Printf("tar:%c\nurl：%s\n\n", targetChar, submitUrl)
		r, err = req.Get(submitUrl, cookie)
		if err != nil {
			log.Fatal(err)
		}
		up = "{\"data\":["
		down = "],\"recordsTotal\""
		jsonData = Util.Between(r.String(), up, down)
		json.Unmarshal([]byte(jsonData), &codeQueryData)

		//获取AC代码并拼接题目结构体
		r, err = req.Get(reqCodeUrl+strconv.Itoa(codeQueryData.RunID), cookie)
		if err != nil {
			log.Fatal(err)
		}

		var reqCodeData Jsonstruct.ReqCodeInfo
		jsonData = r.String()

		content = strings.Replace(jsonData, "\\u003c", "<", -1)
		content = strings.Replace(content, "\\u003e", ">", -1)
		content = strings.Replace(content, "\\u0026", "&", -1)
		content = strings.Replace(content, "\\n", "\\r\\n", -1)
		jsonData = content
		json.Unmarshal([]byte(jsonData), &reqCodeData)
		problemData[i].Code = reqCodeData.Code
		//写代码到文件
		Handlefile.On(username, contestID, problemNum, problemData[i].Code)
		problemNum = i + 1

		outContent += "Title:" + problemData[i].Tag + "-" + problemData[i].Title
		outContent += "Description:\r\n" + problemData[i].Description
		outContent += "\r\nInput:\r\n" + problemData[i].Input
		outContent += "\r\nOutput:\r\n" + problemData[i].Output
		outContent += "\r\nSampleInput:\r\n " + problemData[i].SampleInput
		outContent += "\r\nSampleOutput:\r\n" + problemData[i].SampleOutput + "\r\n\r\n"

		outCode += "Title: " + problemData[i].Tag + "-" + problemData[i].Title
		outCode += "\r\nCode: \r\n" + problemData[i].Code + "\r\n\r\n"

		//写入word
		Handledocx.On(outContent, outCode)
	}
}
