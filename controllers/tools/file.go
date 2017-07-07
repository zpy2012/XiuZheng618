package tools

import (
	"os"
	"strings"
	"mime/multipart"
	"strconv"
	"io"
	"fmt"
)

//保存图片
func SaveImage(file multipart.File, handler *multipart.FileHeader, id int64, filePath string) (path string, err error)  {
	if err != nil {
		return
	}
	defer file.Close()
	if !IsDirExist(filePath) {
		os.MkdirAll(filePath, 0777)
	}
	n := strings.Split(handler.Filename, ".")
	fileName := strconv.FormatInt(id, 10)+"."+n[len(n)-1]
	dir := filePath+fileName
	f, err := os.Create(dir)
	CheckError(err)
	defer f.Close()
	a, err := io.Copy(f, file)
	fmt.Println(a)
	return dir, err
}

//删除图片
func DeleteImageWithPath(path string) (err error) {
	p := strings.Split(path, "/")
	filePath := "."+"/"+p[len(p)-3]+"/"+p[len(p)-2]+"/"+p[len(p)-1]
	if err := os.Remove(filePath);err !=nil {
		return nil
	}else {
		CheckError(err)
		return err
	}
}