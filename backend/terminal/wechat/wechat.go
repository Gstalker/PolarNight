package wechat

import (
	"github.com/gstalker/PolarNight/terminal/types"
	"github.com/gstalker/PolarNight/terminal/wechat/wcf"
)

// NewWechatTerminal 获取微信终端
func NewWechatTerminal() (types.Terminal, error) {
	return wcf.NewWCFClient()
}
