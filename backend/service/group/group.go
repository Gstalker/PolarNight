package group

import (
	"time"

	"github.com/gstalker/PolarNight/terminal"
	"github.com/gstalker/PolarNight/types"
	"github.com/gstalker/PolarNight/utility"
	"github.com/sirupsen/logrus"
)

// ProcessGroupMessage 处理群聊消息
func ProcessGroupMessage(m types.Message) {
	if m == nil {
		logrus.Warn("nil message detected")
		return
	} else if !m.IsGroup() {
		logrus.Warn("not a group message, ignore")
		return
	}
	logrus.Infof("Recieve group message %+v", m)

	if m.RoomID() == "" && m.IsDirectPin() {
		utility.Sleep(time.Second)
		terminal.SendMessage("[呲牙] 被直接at了", "", nil)
	}

	if m.RoomID() == "" && m.IsGlobalPin() {
		utility.Sleep(time.Second)
		terminal.SendMessage("[呲牙] 被全体at了", "", nil)
	}
}
