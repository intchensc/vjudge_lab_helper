package Handlefile

import (
	"fmt"
	"os"
	"strconv"
)

func On(username string, contestNum string, problemNum int, code string) {
	if !isExist("./code") {
		os.Mkdir("./code", os.ModePerm)
	}
	fileName := "./code" + username + contestNum + "_" + strconv.Itoa(problemNum) + ".cpp"
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
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
		fmt.Println(err)
		return false
	}
	return true
}
