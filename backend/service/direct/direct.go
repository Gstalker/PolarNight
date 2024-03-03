package direct

import (
	"github.com/gstalker/PolarNight/types"
	"github.com/sirupsen/logrus"
)

// ProcessDirectMessage 处理私聊消息
func ProcessDirectMessage(m types.Message) {
	if m == nil {
		logrus.Warn("nil message detected")
		return
	} else if m.IsGroup() {
		logrus.Warn("not a direct message, ignore")
		return
	}
	// TODO - finish me
}
