package service

import (
	"math/rand"
	"time"

	"github.com/gstalker/PolarNight/service/direct"
	"github.com/gstalker/PolarNight/service/group"
	"github.com/gstalker/PolarNight/types"
	"github.com/gstalker/PolarNight/utility"
)

// Init 初始化服务
func Init() error {
	return nil
}

// GetMessageProcessor 获取消息处理器
func GetMessageProcessor() func(types.Message) {
	return messageProcessor
}

func messageProcessor(m types.Message) {
	// 添加随机延迟，以避免检测
	utility.Sleep(randomDelay())

	if !m.IsGroup() {
		direct.ProcessDirectMessage(m)
	} else {
		group.ProcessGroupMessage(m)
	}
}

// randomDelay 随机延迟0~1s
func randomDelay() time.Duration {
	return time.Millisecond * time.Duration(rand.Uint32()%1000)
}
