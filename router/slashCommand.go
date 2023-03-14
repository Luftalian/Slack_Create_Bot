// 未実装のため使用不可。CommandBodyの構造体を定義がわからない。
package router

import (
	"fmt"
	"net/http"

	"github.com/Luftalian/Slack_Create_Bot/slackPost"

	"github.com/labstack/echo/v4"
)

type CommandBody struct {
	Token          string `json:"token"`
	TeamID         string `json:"team_id"`
	TeamDomain     string `json:"team_domain"`
	EnterpriseID   string `json:"enterprise_id,omitempty"`
	EnterpriseName string `json:"enterprise_name,omitempty"`
	ChannelID      string `json:"channel_id"`
	ChannelName    string `json:"channel_name"`
	UserID         string `json:"user_id"`
	UserName       string `json:"user_name"`
	Command        string `json:"command"`
	Text           string `json:"text"`
	ResponseURL    string `json:"response_url"`
	TriggerID      string `json:"trigger_id"`
	APIAppID       string `json:"api_app_id"`
}

func HandleCommand(c echo.Context) error {
	fmt.Println("handleCommand")
	// リクエストのボディをJSONとしてパースする
	var response CommandBody
	if err := c.Bind(&response); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// result := map[string]interface{}{
	// 	"text": response.Event.Text,
	// }
	newRequestBody := slackPost.NewRequestBodyFunc("command")

	// POST先のURL
	sendUrl := response.ResponseURL

	err := newRequestBody.PostMessageFunc(sendUrl)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// JSONデータをレスポンスとして返す
	return c.JSON(http.StatusOK, nil)
}
