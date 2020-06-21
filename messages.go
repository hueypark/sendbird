package sendbird

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Message represents a message.
type Message struct {
	User    User   `json:"user"`
	Message string `json:"message"`
}

// Messages returns messages in channel
func (sb *Sendbird) Messages(
	channelType string, channelURL string, unixTimestamp int64,
) ([]Message, error) {
	resp, err := sb.Request(
		http.MethodGet,
		fmt.Sprintf(
			"/%s/%s/messages?message_ts=%v",
			channelType,
			channelURL,
			unixTimestamp*1000,
		),
		nil)
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
	channelType string, channelURL string, userID string, message string,
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
