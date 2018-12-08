package handler

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// File Upload (Put)テスト用
func UploadImage(c echo.Context) error {

	fmt.Println("UploadImage")

	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	fmt.Println("bodyBytes len=", len(bodyBytes))

	_, err := os.Stat("./tmp")
	if err != nil {
		if err := os.Mkdir("./tmp", 0777); err != nil {
			panic(err)
		}
	}

	fileName := getFileName()
	ioutil.WriteFile(fileName, bodyBytes, os.ModePerm)

	mimeType := c.Request().Header.Get("Content-Type")
	fmt.Println("mimeType=", mimeType)

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", fileName))

}

func getFileName() string {
	const format = "20060102_150405" // 24h表現、0埋めあり
	now_date := time.Now().Format(format)
	var buff = bytes.NewBuffer(make([]byte, 0, 100))
	buff.WriteString("./tmp/")
	buff.WriteString(now_date)
	buff.WriteString(".jpg")
	return buff.String()
}
