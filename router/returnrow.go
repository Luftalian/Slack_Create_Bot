package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Luftalian/Slack_Create_Bot/slackPost"

	"github.com/labstack/echo/v4"
)

// type Request struct {
//     // JSONデータを格納するフィールド
//     Data interface{} `json:"data"`
// }

func HandleReturnRaw(c echo.Context) error {
	fmt.Println("handleReturnRaw")
	var req interface{}
	if err := c.Bind(req); err != nil {
		fmt.Println("Bind error")
		fmt.Println(err)
		return err
	}

	// POST先のURL
	sendUrl := os.Getenv("SLACKURL")

	err := slackPost.PostMessageFunc(req, sendUrl)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println("No error")
	return c.JSON(http.StatusOK, req)
}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("/")
//     // POSTリクエストであることを確認する
//     if r.Method != "POST" {
//         http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//         return
//     }
//     // Content-Typeがapplication/jsonであることを確認する
//     if r.Header.Get("Content-Type") != "application/json" {
//         http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
//         return
//     }
// 	fmt.Println(r.Body)
//     // リクエストボディをパースする
//     var reqBody interface{}
//     if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
//         http.Error(w, "Bad request", http.StatusBadRequest)
//         return
//     }

// 	sendUrl := os.Getenv("SLACKURL")
// 	err := slackPost.PostMessageFunc(reqBody, sendUrl)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(reqBody)

//     // パースされたJSONデータをそのまま返す
//     w.Header().Set("Content-Type", "application/json")
//     if err := json.NewEncoder(w).Encode(reqBody); err != nil {
//         log.Println("Failed to encode response:", err)
//     }
// })
// http.HandleFunc("/api/command", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("/api/command")
//     // // POSTリクエストであることを確認する
//     // if r.Method != "POST" {
//     //     http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//     //     return
//     // }
//     // // Content-Typeがapplication/jsonであることを確認する
//     // if r.Header.Get("Content-Type") != "application/json" {
//     //     http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
//     //     return
//     // }
// 	fmt.Println(r.Body)
//     // リクエストボディをパースする
//     var reqBody interface{}
//     if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
//         http.Error(w, "Bad request", http.StatusBadRequest)
//         return
//     }

// 	sendUrl := os.Getenv("SLACKURL")
// 	err := slackPost.PostMessageFunc(reqBody, sendUrl)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(reqBody)

//     // パースされたJSONデータをそのまま返す
//     w.Header().Set("Content-Type", "application/json")
//     if err := json.NewEncoder(w).Encode(reqBody); err != nil {
//         log.Println("Failed to encode response:", err)
//     }
// })
// http.HandleFunc("/api/events", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("/api/events")
//     // POSTリクエストであることを確認する
//     if r.Method != "POST" {
//         http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//         return
//     }
//     // Content-Typeがapplication/jsonであることを確認する
//     if r.Header.Get("Content-Type") != "application/json" {
//         http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
//         return
//     }
//     // リクエストボディをパースする
//     var reqBody interface{}
//     if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
//         http.Error(w, "Bad request", http.StatusBadRequest)
//         return
//     }

// 	sendUrl := os.Getenv("SLACKURL")
// 	err := slackPost.PostMessageFunc(reqBody, sendUrl)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

//     // パースされたJSONデータをそのまま返す
//     w.Header().Set("Content-Type", "application/json")
//     if err := json.NewEncoder(w).Encode(reqBody); err != nil {
//         log.Println("Failed to encode response:", err)
//     }
// })

// // サーバーの開始
// log.Fatal(http.ListenAndServe(":8080", nil))
