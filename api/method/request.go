package method

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var Token string

type Requester interface {
	GetMethodName() string
}

func Request(m Requester) ([]byte, error) {
	if Token == "" {
		return nil, ErrorMsg{Msg:"Token is empty"}
	}

	buf, err := json.Marshal(m)
	if err != nil {
		return  nil, err
	}
	bufferParams := bytes.NewBuffer(buf)

	url := "https://api.telegram.org/bot"
	methodName := "sendMessage"
	fullUrl := url + Token + "/" + methodName

	resp, err := http.Post(fullUrl, "application/json", bufferParams)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrorMsg{Msg:"Some went wrong, raw response=" +  fmt.Sprint(string(bodyBytes))}
	}

	return bodyBytes, err
}
