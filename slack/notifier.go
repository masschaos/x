package slack

import (
	"github.com/slack-go/slack"
)

// Msg is a interface that can get slack msgOption
type Msg interface {
	GetMsgOption() slack.MsgOption
}

// Notifier can send msg to slack
type Notifier struct {
	client  *slack.Client
	channel string
}

// New create a instance of Notifier
func New(token, channel string) Notifier {
	return Notifier{
		client:  slack.New(token),
		channel: channel,
	}
}

// Notify send a message to slack channel
func (n Notifier) Notify(msg Msg) error {
	_, _, err := n.client.PostMessage(n.channel, msg.GetMsgOption())
	return err
}
