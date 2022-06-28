package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"superkoders.com/intl/constants"
)

func BuildIntlPath(srcFolderPath string) string {
	currentDir, err := os.Getwd()
	Check(err)
	intlPath := fmt.Sprint(currentDir + "/" + srcFolderPath + "/" + constants.INTL_FOLDER_PATH)
	return intlPath
}

func CreateFolder(folderPath string) {
	err := os.MkdirAll(folderPath, os.ModePerm)
	Check(err)
}

func WriteToFile(filePath, content string) {
	bytesContent := []byte(content)
	f, err := os.Create(filePath)
	Check(err)
	defer f.Close()
	_, err = f.Write(bytesContent)
	Check(err)
}

func AppendToFile(filePath, content string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)

	if _, err := f.Write([]byte(content)); err != nil {
		Check(err)

	}
	if err := f.Close(); err != nil {
		Check(err)
	}
}

func CreateRootFile(fileName, content string, rootPath string) {
	WriteToFile(rootPath+"/"+fileName, content)
}

func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

func GetPackageJson(rootPath string) *os.File {
	jsonFile, err := os.Open(rootPath + constants.PACKAGE_JSON_PATH)
	Check(err)
	return jsonFile
}

func UnmarshalPackageJson(rootPath string) map[string]interface{} {
	jsonFile := GetPackageJson(rootPath)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	mp := map[string]interface{}{}
	err := json.Unmarshal([]byte(byteValue), &mp)
	Check(err)
	return mp
}
