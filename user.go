package sendbird

import (
	"encoding/json"
	"net/http"
	"strings"
)

// User represents a user.
type User struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
}

// CreateUser creates a user.
func (sb *Sendbird) CreateUser(id string, nickname string) (string, error) {
	body := struct {
		UserID     string `json:"user_id"`
		Nickname   string `json:"nickname"`
		ProfileURL string `json:"profile_url"`
	}{
		UserID:   id,
		Nickname: nickname,
	}

	bs, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	return sb.Request(
		http.MethodPost,
		"/users",
		strings.NewReader(string(bs)),
	)
}
