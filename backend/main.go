package main

import (
	"github.com/gstalker/PolarNight/service"
	"github.com/gstalker/PolarNight/terminal"
	"github.com/gstalker/PolarNight/terminal/types"
	"github.com/sirupsen/logrus"
)

// Init 初始化各个包
func Init() (err error) {
	err = terminal.Init()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := Init()
	if err != nil {
		logrus.WithError(err).Errorf("service init failed, reason: %v", err)
		return
	}

	logrus.Infof("Init finished. current terminal status: %v", terminal.IsLogin())

	err = terminal.SetRecieveStatus(types.RecieveStatusOpen)
	if err != nil {
		logrus.WithError(err).Errorf("enable recieve status failed")
		return
	}

	// 下面这一行是阻塞式的
	terminal.OnMsg(service.GetMessageProcessor())
}
