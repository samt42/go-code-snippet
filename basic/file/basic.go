package main

import (
	"fmt"
)

func main() {
}

func Listfunc(path string, f os.FileInfo, err error) error {
	var strRet string
	strRet, _ = os.Getwd()
	if ostype == "windows" {
		strRet += "\\"
	} else if ostype == "linux" {
		strRet += "/"
	}
	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}
	listfile = append(listfile, path) //将目录push到listfile []string中
	return nil
}

func GetFileList(path string) (string, error) {
	err := filepath.Walk(path, Listfunc)
	if err != nil {
		return "", err
	}
	return " ", nil
}
