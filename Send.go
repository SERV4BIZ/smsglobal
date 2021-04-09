package smsglobal

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Send sms to number phone
func (me *SMSGLOBAL) Send(txtToNumber string, txtMessage string) (string, error) {
	params := url.Values{}
	params.Add("action", "sendsms")
	params.Add("user", me.Username)
	params.Add("password", me.Password)
	params.Add("from", me.SenderName)
	params.Add("to", txtToNumber)
	params.Add("text", txtMessage)
	params.Add("api", "1")

	url := "https://api.smsglobal.com/http-api.php?" + params.Encode()
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	txtBuffer := string(body)
	if strings.HasPrefix(txtBuffer, "OK") {
		splits := strings.Split(txtBuffer, "SMSGlobalMsgID:")
		return splits[1], nil
	}

	return "", errors.New(txtBuffer)
}
