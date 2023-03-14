package databaseConnect

import (
	"database/sql"
	"errors"
	"log"
	"sort"

	"github.com/Luftalian/Slack_Create_Bot/md"
)

type DbRequestBody md.Request

func (r DbRequestBody) ReadMdRequest() ([]int, []string, []string, []string, []string, []string, []string, []string, []string, error) {
	MdRequests := []md.Request{}
	db := Db
	if err := db.Select(&MdRequests, "SELECT * FROM `md`"); errors.Is(err, sql.ErrNoRows) {
		log.Printf("no such md = %s", "___")
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	} else if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	sort.Slice(MdRequests, func(i, j int) bool {
		return MdRequests[i].Id < MdRequests[j].Id
	})
	// log.Printf("MdRequest = %s", MdRequests)
	ids := make([]int, len(MdRequests))
	permissions := make([]string, len(MdRequests))
	titles := make([]string, len(MdRequests))
	contents := make([]string, len(MdRequests))
	commentPermissions := make([]string, len(MdRequests))
	readPermissions := make([]string, len(MdRequests))
	writePermissions := make([]string, len(MdRequests))
	sendTexts := make([]string, len(MdRequests))
	places := make([]string, len(MdRequests))
	for i := 0; i < len(MdRequests); i++ {
		ids[i] = MdRequests[i].Id
		permissions[i] = MdRequests[i].Permission
		titles[i] = MdRequests[i].Title
		contents[i] = MdRequests[i].Content
		commentPermissions[i] = MdRequests[i].CommentPermission
		readPermissions[i] = MdRequests[i].ReadPermission
		writePermissions[i] = MdRequests[i].WritePermission
		sendTexts[i] = MdRequests[i].SendText
		places[i] = MdRequests[i].Place
	}
	return ids, permissions, titles, contents, commentPermissions, readPermissions, writePermissions, sendTexts, places, nil
}

func SetMdRequest(permission string, title string, content string, commentPermission string, sendText string) error {
	db := Db
	if _, err := db.Exec("UPDATE `md` SET `permission` = ?, `title` = ?, `content` = ?, `commentPermission` = ?, `sendText` = ?", permission, title, content, commentPermission, sendText); err != nil {
		return err
	}
	return nil
}

func UpdatePermission(id int, permission string) error {
	db := Db
	if _, err := db.Exec("UPDATE `md` SET `permission` = ? WHERE `id` = ?", permission, id); err != nil {
		return err
	}
	return nil
}

func UpdateTitle(id int, title string) error {
	db := Db
	if _, err := db.Exec("UPDATE `md` SET `title` = ? WHERE `id` = ?", title, id); err != nil {
		return err
	}
	return nil
}

func UpdateContent(id int, content string) error {
	db := Db
	if _, err := db.Exec("UPDATE `md` SET `content` = ? WHERE `id` = ?", content, id); err != nil {
		return err
	}
	return nil
}

func UpdateCommentPermission(id int, commentPermission string) error {
	db := Db
	if _, err := db.Exec("UPDATE `md` SET `commentPermission` = ? WHERE `id` = ?", commentPermission, id); err != nil {
		return err
	}
	return nil
}

func UpdateReadPermission(id int, readPermission string) error {
	db := Db
	if _, err := db.Exec("UPDATE `md` SET `readPermission` = ? WHERE `id` = ?", readPermission, id); err != nil {
		return err
	}
	return nil
}

func UpdateWritePermission(id int, writePermission string) error {
	db := Db
	if _, err := db.Exec("UPDATE `md` SET `writePermission` = ? WHERE `id` = ?", writePermission, id); err != nil {
		return err
	}
	return nil
}

func UpdateSendText(id int, sendText string) error {
	db := Db
	if _, err := db.Exec("UPDATE `md` SET `sendText` = ? WHERE `id` = ?", sendText, id); err != nil {
		return err
	}
	return nil
}

func UpdatePlace(id int, place string) error {
	db := Db
	if _, err := db.Exec("UPDATE `md` SET `place` = ? WHERE `id` = ?", place, id); err != nil {
		return err
	}
	return nil
}
