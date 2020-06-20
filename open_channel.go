package sendbird

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// OpenChannels returns open channels.
func (sb *Sendbird) OpenChannels() (string, error) {
	return sb.Request(http.MethodGet, fmt.Sprintf("/%s", OpenChannels), nil)
}

// CreateOpenChannel creates open channel.
func (sb *Sendbird) CreateOpenChannel(name string, channelURL string) (string, error) {
	body := struct {
		Name       string `json:"name"`
		ChannelURL string `json:"channel_url"`
	}{
		Name:       name,
		ChannelURL: channelURL,
	}

	bs, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	return sb.Request(
		http.MethodPost,
		fmt.Sprintf("/%s", OpenChannels),
		strings.NewReader(string(bs)),
	)
}
