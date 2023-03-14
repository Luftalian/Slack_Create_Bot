package cronAction

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Luftalian/Slack_Create_Bot/databaseConnect"
	"github.com/Luftalian/Slack_Create_Bot/md"
	"github.com/Luftalian/Slack_Create_Bot/slackPost"
	"github.com/Luftalian/Slack_Create_Bot/timeSet"
)

type readTimeFunc interface {
	ReadDateHour() ([]int, []int, []int, []int, error)
}

func GetTimeFunc(r readTimeFunc) ([]int, []int, []int, []int, error) {
	ids, dates, hours, minutes, err := r.ReadDateHour()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return ids, dates, hours, minutes, nil
}

type readMdFunc interface {
	ReadMdRequest() ([]int, []string, []string, []string, []string, []string, []string, []string, []string, error)
}

func GetMdFunc(r readMdFunc) ([]int, []string, []string, []string, []string, []string, []string, []string, []string, error) {
	ids, permissions, titles, contents, commentPermissions, readPermissions, writePermissions, sendTexts, places, err := r.ReadMdRequest()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	return ids, permissions, titles, contents, commentPermissions, readPermissions, writePermissions, sendTexts, places, nil
}

func CronActionFunc() {
	fmt.Println("cronActionFunc")
	ids, dates, hours, minutes, err := GetTimeFunc(&databaseConnect.TimeDataBase{})
	if err != nil {
		fmt.Println(err)
	}
	ids_, permissions, titles, contents, commentPermissions, readPermissions, writePermissions, sendTexts, places, err := GetMdFunc(&databaseConnect.DbRequestBody{})
	if err != nil {
		fmt.Println(err)
	}

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err)
	}
	t := time.Now().In(loc)

	yearString := strconv.Itoa(t.Year())
	monthString := strconv.Itoa(int(t.Month()))
	if len(monthString) == 1 {
		monthString = "0" + monthString
	}

	todayString := yearString[2:4] + "/" + monthString + "/" + strconv.Itoa(t.Day()) + " (" + timeSet.WeekdayFunc(t.Weekday()) + ")"
	fmt.Printf("ids length: %v", len(ids))
	for i := 0; i < len(ids); i++ {
		tomorrowWeekday := int(t.Weekday()) + 1
		if tomorrowWeekday == 7 {
			tomorrowWeekday = 0
		}
		if t.Hour() == hours[i] && t.Minute() == minutes[i] && tomorrowWeekday == dates[i] {
			// Bearerトークン
			accessToken := os.Getenv("TOKEN")

			if ids[i] != ids_[i] {
				fmt.Println("id is not match")
				break
			}
			replaceTime := strings.Replace(contents[i], "[yy/tmp]", todayString, -1)
			r := md.Request{
				Permission:        permissions[i],
				Title:             titles[i],
				Content:           replaceTime,
				CommentPermission: commentPermissions[i],
				ReadPermission:    readPermissions[i],
				WritePermission:   writePermissions[i],
			}

			url, err := r.MakeNewMd(accessToken)
			if err != nil {
				fmt.Println(err)
			}
			minuteString := strconv.Itoa(minutes[i])
			if len(minuteString) == 1 {
				minuteString = "0" + minuteString
			}
			replacedString := strings.Replace(sendTexts[i], "[場所]", places[i], -1)
			replacedString = strings.Replace(replacedString, "[時間]", strconv.Itoa(hours[i])+":"+minuteString, -1)
			replacedString = strings.Replace(replacedString, "[リンク]", url, -1)
			newRequestBody := slackPost.NewRequestBodyFunc(replacedString)

			// POST先のURL
			sendUrl := os.Getenv("SLACKURL")

			err = newRequestBody.PostMessageFunc(sendUrl)
			if err != nil {
				fmt.Println(err)
			}
		}
		if t.Hour() == hours[i] && t.Minute() == minutes[i] && int(t.Weekday()) == dates[i] {
			newRequestBody := slackPost.NewRequestBodyFunc("<!channel> 電装班MTGの開始時刻です")
			// POST先のURL
			sendUrl := os.Getenv("SLACKURL")

			err = newRequestBody.PostMessageFunc(sendUrl)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func CronActionFuncNoTime() {
	fmt.Println("cronActionFuncNoTime")
	ids, _, hours, minutes, err := GetTimeFunc(&databaseConnect.TimeDataBase{})
	if err != nil {
		fmt.Println(err)
	}
	ids_, permissions, titles, contents, commentPermissions, readPermissions, writePermissions, sendTexts, places, err := GetMdFunc(&databaseConnect.DbRequestBody{})
	if err != nil {
		fmt.Println(err)
	}

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err)
	}
	t := time.Now().In(loc)

	yearString := strconv.Itoa(t.Year())
	monthString := strconv.Itoa(int(t.Month()))
	if len(monthString) == 1 {
		monthString = "0" + monthString
	}

	todayString := yearString[2:4] + "/" + monthString + "/" + strconv.Itoa(t.Day()) + " (" + timeSet.WeekdayFunc(t.Weekday()) + ")"

	fmt.Printf("ids length: %v", len(ids))
	for i := 0; i < len(ids); i++ {
		// if t.Hour() == hours[i] && t.Minute() == minutes[i] && int(t.Weekday()) == dates[i] {
		// Bearerトークン
		accessToken := os.Getenv("TOKEN")

		if ids[i] != ids_[i] {
			fmt.Println("id is not match")
			break
		}
		replaceTime := strings.Replace(contents[i], "[yy/tmp]", todayString, -1)
		r := md.Request{
			Permission:        permissions[i],
			Title:             titles[i],
			Content:           replaceTime,
			CommentPermission: commentPermissions[i],
			ReadPermission:    readPermissions[i],
			WritePermission:   writePermissions[i],
		}

		url, err := r.MakeNewMd(accessToken)
		if err != nil {
			fmt.Println(err)
		}
		minuteString := strconv.Itoa(minutes[i])
		if len(minuteString) == 1 {
			minuteString = "0" + minuteString
		}
		replacedString := strings.Replace(sendTexts[i], "[場所]", places[i], -1)
		replacedString = strings.Replace(replacedString, "[時間]", strconv.Itoa(hours[i])+":"+minuteString, -1)
		replacedString = strings.Replace(replacedString, "[リンク]", url, -1)
		newRequestBody := slackPost.NewRequestBodyFunc(replacedString)

		// POST先のURL
		sendUrl := os.Getenv("SLACKURL")

		err = newRequestBody.PostMessageFunc(sendUrl)
		if err != nil {
			fmt.Println(err)
		}
		// }
	}
}
