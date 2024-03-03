package terminal

import (
	"fmt"

	"github.com/gstalker/PolarNight/config"
	"github.com/gstalker/PolarNight/terminal/types"
	"github.com/gstalker/PolarNight/terminal/wechat"
)

// messageTerminal 聊天信息收发终端
var messageTerminal types.Terminal

func Init() (err error) {
	switch config.GetTerminalType() {
	case config.TerminalTypeWechat:
		messageTerminal, err = wechat.NewWechatTerminal()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf(types.ErrInvalidTerminalType)
	}
	return nil
}

// IsLogin 当前终端账户是否在线
func IsLogin() bool {
	return messageTerminal.IsLogin()
}

// OnMsg 回调注册器：消息处理函数
// 注意：该函数会阻塞当前goroutine
func OnMsg(f func(types.Message)) error {
	return messageTerminal.OnMsg(f)
}

// SendMessage 发送信息
// content: 消息正文
// id: 消息接收者id，群聊或者个人用户
// atGroup: @列表，个人用户
func SendMessage(content types.MessageContent, id types.EntityID, atGroup []types.EntityID) {
	messageTerminal.SendMessage(content, id, atGroup)
}

// SetRecieveStatus 设置消息接收状态
func SetRecieveStatus(status types.RecieveStatus) error {
	return messageTerminal.SetRecieveStatus(status)
}
