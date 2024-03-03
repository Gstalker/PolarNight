package types

// MessageContent 消息正文
type MessageContent string

// EntityID 用户/群聊ID
type EntityID string

// MessageID 消息ID
type MessageID uint64

// RecieveStatus 消息接收开关状态
type RecieveStatus bool

const (
	// RecieveStatusOpen 开启消息接收
	RecieveStatusOpen = true

	// RecieveStatusClose 关闭消息接接收
	RecieveStatusClose = false
)

// Message 聊天信息接口
// 由于各个平台的信息格式不一样，因此抽象该接口出来，用于统一对接各个平台的消息结构体
type Message interface {
	// IsSelf 该消息由当前终端账户发出
	IsSelf() bool

	// IsGroup 该消息来自群组消息，而非私聊消息
	IsGroup() bool

	// IsDirectPin 仅群聊中，如果被直接at（pin），则返回true
	IsDirectPin() bool

	// IsGlobalPin 仅群聊中，如果存在@所有人；或者类似discord的@group，且当前终端账户在该group中，则返回true
	IsGlobalPin() bool

	// MessageId 当前消息ID
	MessageId() MessageID

	// RoomID 如果该消息来自群组消息，则该字段有效，代表该消息所属的群组的id
	RoomID() EntityID

	// Content 消息正文
	Content() string

	// Sender 消息发送者的id
	Sender() EntityID
}

// Terminal 聊天程序终端抽象接口
type Terminal interface {
	// IsLogin 当前终端账户是否在线
	IsLogin() bool

	// OnMsg 回调注册器：消息处理函数
	// 注意：该函数会阻塞当前goroutine
	OnMsg(func(Message)) error

	// SendMessage 发送信息
	// content: 消息正文
	// id: 消息接收者id，群聊或者个人用户
	// atGroup: @列表，个人用户
	SendMessage(content MessageContent, id EntityID, atGroup []EntityID)

	// SetRecieveStatus 设置消息接收状态
	SetRecieveStatus(status RecieveStatus) error
}
