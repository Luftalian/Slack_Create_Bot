package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Authorization struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

func HandleRequestAndReturnJSON(c echo.Context) error {
	fmt.Println("handleRequestAndReturnJSON")
	// リクエストのボディをJSONとしてパースする
	var auth Authorization
	if err := c.Bind(&auth); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	result := map[string]interface{}{
		"challenge": auth.Challenge,
	}

	// JSONデータをレスポンスとして返す
	return c.JSON(http.StatusOK, result)
}
