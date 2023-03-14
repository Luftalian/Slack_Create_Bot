package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Luftalian/Slack_Create_Bot/brain"
	"github.com/Luftalian/Slack_Create_Bot/slackPost"

	"github.com/labstack/echo/v4"
)

type Response1 struct {
	APIAppID       string `json:"api_app_id"`
	Authorizations []struct {
		EnterpriseID        interface{} `json:"enterprise_id"`
		IsBot               bool        `json:"is_bot"`
		IsEnterpriseInstall bool        `json:"is_enterprise_install"`
		TeamID              string      `json:"team_id"`
		UserID              string      `json:"user_id"`
	} `json:"authorizations"`
	ContextEnterpriseID interface{} `json:"context_enterprise_id"`
	ContextTeamID       string      `json:"context_team_id"`
	Event               struct {
		Blocks []struct {
			BlockID  string `json:"block_id"`
			Elements []struct {
				Elements []struct {
					Text string `json:"text"`
					Type string `json:"type"`
				} `json:"elements"`
				Type string `json:"type"`
			} `json:"elements"`
			Type string `json:"type"`
		} `json:"blocks"`
		Channel     string `json:"channel"`
		ChannelType string `json:"channel_type"`
		ClientMsgID string `json:"client_msg_id"`
		EventTS     string `json:"event_ts"`
		Team        string `json:"team"`
		Text        string `json:"text"`
		Ts          string `json:"ts"`
		Type        string `json:"type"`
		User        string `json:"user"`
	} `json:"event"`
	EventContext       string  `json:"event_context"`
	EventID            string  `json:"event_id"`
	EventTime          float64 `json:"event_time"`
	IsExtSharedChannel bool    `json:"is_ext_shared_channel"`
	TeamID             string  `json:"team_id"`
	Token              string  `json:"token"`
	Type               string  `json:"type"`
}

func HandlePostMessageJSON(c echo.Context) error {
	fmt.Println("handlePostMessageJSON")
	// リクエストのボディをJSONとしてパースする
	var response Response1
	if err := c.Bind(&response); err != nil {
		return c.NoContent(http.StatusOK)
	}
	// fmt.Println("=====================================")
	// fmt.Println(response.APIAppID)
	// fmt.Println(response.Authorizations)
	// for _, auth := range response.Authorizations {
	// 	fmt.Println(auth.EnterpriseID)
	// 	fmt.Println(auth.IsEnterpriseInstall)
	// 	fmt.Println(auth.IsBot)
	// 	fmt.Println(auth.TeamID)
	// 	fmt.Println(auth.UserID)
	// }
	// fmt.Println(response.ContextEnterpriseID)
	// fmt.Println(response.ContextTeamID)
	// fmt.Println(response.Event.Blocks)
	// for _, block := range response.Event.Blocks {
	// 	fmt.Println(block.BlockID)
	// 	fmt.Println(block.Elements)
	// 	for _, element := range block.Elements {
	// 		fmt.Println(element.Elements)
	// 		for _, element2 := range element.Elements {
	// 			fmt.Println(element2.Text)
	// 			fmt.Println(element2.Type)
	// 		}
	// 		fmt.Println(element.Type)
	// 	}
	// 	fmt.Println(block.Type)
	// }
	// fmt.Println(response.Event.Channel)
	// fmt.Println(response.Event.ChannelType)
	// fmt.Println(response.Event.ClientMsgID)
	// fmt.Println(response.Event.EventTS)
	// fmt.Println(response.Event.Team)
	// fmt.Println(response.Event.Text)
	// fmt.Println(response.Event.Ts)
	// fmt.Println(response.Event.Type)
	// fmt.Println(response.Event.User)
	// fmt.Println(response.EventContext)
	// fmt.Println(response.EventID)
	// fmt.Println(response.EventTime)
	// fmt.Println(response.IsExtSharedChannel)
	// fmt.Println(response.TeamID)
	// fmt.Println(response.Token)
	// fmt.Println(response.Type)
	// fmt.Println("=====================================")

	if !response.Authorizations[0].IsBot {

		aaa, err := brain.CheckWords(response.Event.Text)
		if err != nil {
			return c.NoContent(http.StatusOK)
		}
		if aaa == "" {
			return c.NoContent(http.StatusOK)
		}

		newRequestBody := slackPost.NewRequestBodyFunc(aaa)

		// POST先のURL
		sendUrl := os.Getenv("SLACKURL")

		err = newRequestBody.PostMessageFunc(sendUrl)
		if err != nil {
			return c.NoContent(http.StatusOK)
		}

		// JSONデータをレスポンスとして返す
		return c.NoContent(http.StatusOK)
	}
	fmt.Println("From Bot")
	return c.NoContent(http.StatusOK)
}
