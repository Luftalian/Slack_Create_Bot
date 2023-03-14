package md

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Request struct {
	Id                int    `db:"id"`
	Permission        string `db:"permission"`
	Title             string `db:"title"`
	Content           string `db:"content"`
	CommentPermission string `db:"commentPermission"`
	ReadPermission    string `db:"readPermission"`
	WritePermission   string `db:"writePermission"`
	SendText          string `db:"sendText"`
	Place             string `db:"place"`
}

type Response2 struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Tags           []any  `json:"tags"`
	CreatedAt      int64  `json:"createdAt"`
	PublishType    string `json:"publishType"`
	PublishedAt    any    `json:"publishedAt"`
	Permalink      any    `json:"permalink"`
	PublishLink    string `json:"publishLink"`
	ShortID        string `json:"shortId"`
	Content        string `json:"content"`
	LastChangedAt  int64  `json:"lastChangedAt"`
	LastChangeUser struct {
		Name      string `json:"name"`
		Photo     string `json:"photo"`
		Biography any    `json:"biography"`
		UserPath  string `json:"userPath"`
	} `json:"lastChangeUser"`
	UserPath        any    `json:"userPath"`
	TeamPath        string `json:"teamPath"`
	ReadPermission  string `json:"readPermission"`
	WritePermission string `json:"writePermission"`
}

func MakeNewRequestBody(permission string, title string, content string, commentPermission string, readPermission string, writePermission string) Request {
	return Request{
		Permission:        permission,
		Title:             title,
		Content:           content,
		CommentPermission: commentPermission,
		ReadPermission:    readPermission,
		WritePermission:   writePermission,
	}
}

func (r Request) MakeNewMd(accessToken string) (string, error) {
	// HTTP POSTリクエストの送信
	requestBody, err := json.Marshal(map[string]string{
		"permission":        r.Permission,
		"title":             r.Title,
		"content":           r.Content,
		"commentPermission": r.CommentPermission,
		"readPermission":    r.ReadPermission,
		"writePermission":   r.WritePermission,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, "https://api.hackmd.io/v1/teams/"+os.Getenv("TEAMID")+"/notes", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	// HTTP POSTリクエストの送信
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	// レスポンスの読み取り
	var response2 Response2
	err = json.NewDecoder(resp.Body).Decode(&response2)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return "https://hackmd.io/" + response2.ID, nil
}
