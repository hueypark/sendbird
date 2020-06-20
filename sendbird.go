package sendbird

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Sendbird represents sendbird client.
type Sendbird struct {
	applicationID string
	apiToken      string
}

// NewSendbird creates new sendbird.
func NewSendbird(applicationID string, apiToken string) (*Sendbird, error) {
	sb := &Sendbird{
		applicationID: applicationID,
		apiToken:      apiToken,
	}

	return sb, nil

}

func (sb *Sendbird) Request(method string, endpoint string, body io.Reader) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("https://api-%s.sendbird.com/v3%s", sb.applicationID, endpoint),
		body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json, charset=utf8")
	req.Header.Add("Api-Token", sb.apiToken)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf(
			"status: %s, method: %s, reqURL: %s, respBody: %s",
			resp.Status,
			req.Method,
			req.URL,
			string(respBody),
		)
	}

	return string(respBody), nil

}
