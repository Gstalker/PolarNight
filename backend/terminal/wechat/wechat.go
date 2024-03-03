package wechat

import (
	"github.com/gstalker/PolarNight/terminal/wechat/wcf"
	"github.com/gstalker/PolarNight/types"
)

// NewWechatTerminal 获取微信终端
func NewWechatTerminal() (types.Terminal, error) {
	return wcf.NewWCFClient()
}
