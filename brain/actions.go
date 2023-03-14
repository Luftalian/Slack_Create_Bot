package brain

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Luftalian/Slack_Create_Bot/cronAction"
	"github.com/Luftalian/Slack_Create_Bot/databaseConnect"
	"github.com/Luftalian/Slack_Create_Bot/md"
)

var detailString string = "```!md-bot-help: ヘルプを表示します。\n!md-bot-set [id] time [hour] [minute]: 指定したidのmdを指定した時間に送信します。\n!md-bot-md: 新しいmdを作成します。\n!md-bot-set [id] [permission, title, content, commentPermission, sendText, date, place] [value]: 値を変更```\n\n詳しくはhttps://hackmd.io/711nNKo4RmeyZEOMWcKdEg?view"

func CheckWords(words string) (string, error) {
	if strings.HasPrefix(words, "!md-bot-help") {
		return detailString, nil
	}
	if strings.HasPrefix(words, "!md-bot-md") {
		accessToken := os.Getenv("TOKEN")
		r := md.Request{
			Permission:        "freely",
			Title:             "Title",
			Content:           "",
			CommentPermission: "everyone",
			ReadPermission:    "guest",
			WritePermission:   "guest",
		}
		url, err := r.MakeNewMd(accessToken)
		if err != nil {
			fmt.Println(err)
		}
		return url, nil
	}
	if strings.HasPrefix(words, "!md-bot-set") {
		// make new group
		arr1 := strings.Split(words, " ")
		// fmt.Println(len(arr1))
		count := 0
		for _, arr := range arr1 {
			if arr != "" {
				arr1[count] = arr
				count++
			}
		}
		id := arr1[1]
		idInt, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		if count == 5 && arr1[2] == "time" {
			hour, err := strconv.Atoi(arr1[3])
			if err != nil {
				fmt.Println(err)
				return "", err
			}
			minute, err := strconv.Atoi(arr1[4])
			if err != nil {
				fmt.Println(err)
				return "", err
			}
			if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
				return "", fmt.Errorf("invalid time")
			}
			if minute != 0 && minute != 30 {
				return "", fmt.Errorf("invalid time")
			}
			err = databaseConnect.UpdateTime(idInt, hour, minute)
			if err != nil {
				return "", err
			}
			return "success", nil
		}
		if count != 4 {
			return "", fmt.Errorf("invalid command")
		}
		switch arr1[2] {
		case "permission":
			// set permission
			if !md.CheckPermission(arr1[3]) {
				return "", fmt.Errorf("invalid permission")
			}
			err := databaseConnect.UpdatePermission(idInt, arr1[3])
			if err != nil {
				return "", err
			}
		case "title":
			// set title
			err := databaseConnect.UpdateTitle(idInt, arr1[3])
			if err != nil {
				return "", err
			}
		case "content":
			// set content
			err := databaseConnect.UpdateContent(idInt, arr1[3])
			if err != nil {
				return "", err
			}
		case "commentPermission":
			// set comment
			if !md.CheckCommentPermission(arr1[3]) {
				return "", fmt.Errorf("invalid permission")
			}
			err := databaseConnect.UpdateCommentPermission(idInt, arr1[3])
			if err != nil {
				return "", err
			}
		case "readPermission":
			// set read
			if !md.CheckReadPermission(arr1[3]) {
				return "", fmt.Errorf("invalid permission")
			}
			err := databaseConnect.UpdateReadPermission(idInt, arr1[3])
			if err != nil {
				return "", err
			}
		case "writePermission":
			// set write
			if !md.CheckWritePermission(arr1[3]) {
				return "", fmt.Errorf("invalid permission")
			}
			err := databaseConnect.UpdateWritePermission(idInt, arr1[3])
			if err != nil {
				return "", err
			}
		case "sendText":
			// set send
			err := databaseConnect.UpdateSendText(idInt, arr1[3])
			if err != nil {
				return "", err
			}
		case "date":
			// set date
			date, err := strconv.Atoi(arr1[3])
			if err != nil {
				return "", err
			}
			if date < 0 || date > 7 {
				return "", fmt.Errorf("invalid date")
			}
			err = databaseConnect.UpdateDate(idInt, date)
			if err != nil {
				return "", err
			}
		case "place":
			// set place
			err := databaseConnect.UpdatePlace(idInt, arr1[3])
			if err != nil {
				return "", err
			}
		default:
			return "", fmt.Errorf("invalid command")
		}
		return "success", nil
	}
	if strings.HasPrefix(words, "!md-bot-test") {
		// edit request
		cronAction.CronActionFuncNoTime()
		return "", nil
	}
	// if strings.HasPrefix(words, "!md-bot-new") {
	// 	// edit time
	// 	return "ID: "+, nil
	// }
	return "", fmt.Errorf("no command")
}
