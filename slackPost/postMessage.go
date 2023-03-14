package slackPost

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (r RequestBody) PostMessageFunc(sendUrl string) error {
	// POSTリクエストの生成
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, sendUrl, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	// POSTリクエストの送信
	client := http.DefaultClient
	response3, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response3.Body.Close()

	// レスポンスの読み取り
	var responseBody map[string]interface{}
	err = json.NewDecoder(response3.Body).Decode(&responseBody)
	if err != nil {
		return err
	}
	// fmt.Println(responseBody)
	return nil
}

func PostMessageFunc(r interface{}, sendUrl string) error {
	// POSTリクエストの生成
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, sendUrl, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	// POSTリクエストの送信
	client := http.DefaultClient
	response3, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response3.Body.Close()

	// レスポンスの読み取り
	var responseBody map[string]interface{}
	err = json.NewDecoder(response3.Body).Decode(&responseBody)
	if err != nil {
		return err
	}
	// fmt.Println(responseBody)
	return nil
}
