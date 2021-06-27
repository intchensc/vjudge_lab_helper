package handle

import (
	"log"
	"os"
	"strconv"

	"github.com/intxiaoquan/vjudge_lab_helper/jsonstruct"
	"github.com/nguyenthenguyen/docx"
)

func DocOn(outData [20]jsonstruct.Output2File, len int) {
	rr, err := docx.ReadDocxFile("./tmpl.docx")
	if err != nil {
		panic(err)
	}

	docx1 := rr.Editable()
	for i := 0; i < len; i++ {
		docx1.Replace("{{content}}", outData[i].Content, 1)
		docx1.Replace("{{code}}", outData[i].Code, 1)
	}
	docx1.WriteToFile("./vlh_out.docx")
	rr.Close()
}

func FileOn(username string, contestNum int, problemNum int, code string) {
	if !isExist("./code") {
		os.Mkdir("./code", os.ModePerm)
	}
	fileName := "./code/" + username + "_" + strconv.Itoa(contestNum+1) + "_" + strconv.Itoa(problemNum) + ".cpp"
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		log.Println(err.Error())
	} else {
		_, err = f.Write([]byte(code))
	}
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		log.Println(err)
		return false
	}
	return true
}
