package sendbird

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Message represents a message.
type Message struct {
	User      User   `json:"user"`
	Message   string `json:"message"`
	MessageID int64  `json:"message_id"`
}

// MessageReadType is typeof message read.(timestamp or id)
type MessageReadType int

const (
	MessageReadID MessageReadType = iota
	MessageReadTimestamp
)

// Messages returns messages in channel
func (sb *Sendbird) Messages(
	channelType ChannelType,
	channelURL string,
	messageReadType MessageReadType,
	messageReadVal int64,
) ([]Message, error) {
	url := fmt.Sprintf("/%s/%s/messages?", channelType, channelURL)

	switch messageReadType {
	case MessageReadID:
		url += fmt.Sprintf("message_id=%v", messageReadVal)
	case MessageReadTimestamp:
		url += fmt.Sprintf("message_ts=%v", messageReadVal)
	}

	resp, err := sb.Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	j := struct {
		Messages []Message `json:"messages"`
	}{}
	err = json.Unmarshal([]byte(resp), &j)
	if err != nil {
		return nil, err
	}

	return j.Messages, nil
}

// SendMessage sends message.
func (sb *Sendbird) SendMessage(
	channelType ChannelType, channelURL string, userID string, message string,
) (string, error) {
	body := struct {
		MessageType string `json:"message_type"`
		UserID      string `json:"user_id"`
		Message     string `json:"message"`
	}{
		MessageType: "MESG",
		UserID:      userID,
		Message:     message,
	}

	bs, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	return sb.Request(
		http.MethodPost,
		fmt.Sprintf(
			"/%s/%s/messages",
			channelType,
			channelURL,
		),
		strings.NewReader(string(bs)),
	)
}
