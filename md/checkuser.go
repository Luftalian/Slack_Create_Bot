package md

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserPath string `json:"userPath"`
	Photo    string `json:"photo"`
	Teams    []struct {
		ID          string `json:"id"`
		OwnerID     string `json:"ownerId"`
		Path        string `json:"path"`
		Name        string `json:"name"`
		Logo        string `json:"logo"`
		Description any    `json:"description"`
		Visibility  string `json:"visibility"`
		CreatedAt   int64  `json:"createdAt"`
	} `json:"teams"`
}

func CheckEmail(accessToken string) (string, error) {
	// HTTP GETリクエストの送信
	resp, err := http.Get("https://api.hackmd.io/v1/me")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	// AuthorizationヘッダーにBearerトークンを設定
	req, err := http.NewRequest(http.MethodGet, "https://api.hackmd.io/v1/me", nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// HTTP GETリクエストの送信
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	// レスポンスの読み取り
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return response.Email, nil
}
