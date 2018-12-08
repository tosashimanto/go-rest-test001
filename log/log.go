package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var AppLog *logrus.Logger = logrus.New()

func Init() {
	fmt.Println("Log初期化開始")
	file, err := os.OpenFile("controller.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}

	AppLog.Out = file
	AppLog.Formatter = &logrus.JSONFormatter{}
	fmt.Println("Log初期化完了")
}
