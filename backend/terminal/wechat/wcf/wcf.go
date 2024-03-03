package wcf

import (
	"fmt"

	"github.com/danbai225/WeChatFerry-go/wcf"
	"github.com/gstalker/PolarNight/types"
	"github.com/sirupsen/logrus"
)

// wCFClientWrapper WCF客户端包装器
type wCFClientWrapper struct {
	client *wcf.Client
}

// NewWCFClient 获取WCF客户端
// 请参阅：https://github.com/lich0821/wcf-client-rust?tab=readme-ov-file#%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B
func NewWCFClient() (types.Terminal, error) {
	wcfClient, err := wcf.NewWCF("")
	if err != nil {
		return nil, err
	}

	logrus.Infof("WCFClient created")
	logrus.Infof("Current account: %v", wcfClient.GetSelfWXID())

	err = initMessage(wcfClient)
	if err != nil {
		return nil, err
	}

	return &wCFClientWrapper{
		client: wcfClient,
	}, nil
}

// IsLogin 当前终端账户是否在线
func (w *wCFClientWrapper) IsLogin() bool {
	return w.client.IsLogin()
}

// OnMsg 回调注册器：消息处理函数
// 注意：该函数会阻塞当前goroutine
func (w *wCFClientWrapper) OnMsg(f func(types.Message)) error {
	return w.client.OnMSG(func(msg *wcf.WxMsg) {
		f(&wCFMessageWrapper{msg: msg})
	})
}

// SendMessage 发送信息
// content: 消息正文
// id: 消息接收者id，群聊或者个人用户
// atGroup: @列表，个人用户
func (w *wCFClientWrapper) SendMessage(content types.MessageContent, id types.EntityID, atGroup []types.EntityID) {
	atg := []string{}
	for _, i := range atGroup {
		atg = append(atg, string(i))
	}
	w.client.SendTxt(string(content), string(id), atg)
}

// SetRecieveStatus 设置消息接收状态
func (w *wCFClientWrapper) SetRecieveStatus(status types.RecieveStatus) error {
	var result int32
	if status == types.RecieveStatusOpen {
		result = w.client.EnableRecvTxt()
	} else {
		result = w.client.DisableRecvTxt()
	}
	if result != 0 {
		return fmt.Errorf(types.ErrSetRecieveStatusFailed)
	}
	return nil
}
