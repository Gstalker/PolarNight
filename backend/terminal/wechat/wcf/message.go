package wcf

import (
	"fmt"
	"regexp"

	"github.com/danbai225/WeChatFerry-go/wcf"
	"github.com/gstalker/PolarNight/types"
)

const (
	regexpGlobalPin         = "@(?:所有人|all|All)"
	regexpTemplateDirectPin = "<atuserlist>.*(%v).*</atuserlist>"
)

var regexpDirectPin string

func initMessage(c *wcf.Client) error {
	regexpDirectPin = fmt.Sprintf(regexpTemplateDirectPin, c.GetSelfWXID())
	return nil
}

// wCFMessageWrapper WCF微信聊天消息包装器
type wCFMessageWrapper struct {
	msg *wcf.WxMsg
}

// IsSelf 该消息由当前终端账户发出
func (w *wCFMessageWrapper) IsSelf() bool {
	return w.msg.IsSelf
}

// IsGroup 该消息来自群组消息，而非私聊消息
func (w *wCFMessageWrapper) IsGroup() bool {
	return w.msg.IsGroup
}

// IsDirectPin 仅群聊中，如果被直接at（pin），则返回true
func (w *wCFMessageWrapper) IsDirectPin() bool {
	re, err := regexp.Compile(regexpDirectPin)
	if err != nil {
		panic(fmt.Sprintf("not a valid regexp: %v", regexpDirectPin))
	}
	result := re.FindAll([]byte(w.msg.Xml), -1)
	return result != nil && !w.IsGlobalPin()
}

// IsGlobalPin 仅群聊中，如果存在@所有人；或者类似discord的@group，且当前终端账户在该group中，则返回true
func (w *wCFMessageWrapper) IsGlobalPin() bool {
	re, err := regexp.Compile(regexpGlobalPin)
	if err != nil {
		panic(fmt.Sprintf("not a valid regexp: %v", regexpGlobalPin))
	}
	result := re.FindAll([]byte(w.Content()), -1)
	return result != nil
}

// MessageId 当前消息ID
func (w *wCFMessageWrapper) MessageId() types.MessageID {
	return types.MessageID(w.msg.Id)
}

// RoomID 如果该消息来自群组消息，则该字段有效，代表该消息所属的群组的id
func (w *wCFMessageWrapper) RoomID() types.EntityID {
	return types.EntityID(w.msg.Roomid)
}

// Content 消息正文
func (w *wCFMessageWrapper) Content() string {
	return w.msg.Content
}

// Sender 消息发送者的id
func (w *wCFMessageWrapper) Sender() types.EntityID {
	return types.EntityID(w.msg.Sender)
}
