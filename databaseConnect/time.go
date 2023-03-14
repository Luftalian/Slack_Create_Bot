package databaseConnect

import (
	"database/sql"
	"errors"
	"log"
	"sort"
)

type TimeDataBase struct {
	Id     int `db:"id"`
	Date   int `db:"date"` // 0:Sunday, 1:Monday, 2:Tuesday, 3:Wednesday, 4:Thursday, 5:Friday, 6:Saturday
	Hour   int `db:"hour"`
	Minute int `db:"minute"`
}

func (r TimeDataBase) ReadDateHour() ([]int, []int, []int, []int, error) {
	TimeDataBases := []TimeDataBase{}
	db := Db
	if err := db.Select(&TimeDataBases, "SELECT * FROM `time`"); errors.Is(err, sql.ErrNoRows) {
		log.Printf("no such time = %s", "___")
		return nil, nil, nil, nil, err
	} else if err != nil {
		return nil, nil, nil, nil, err
	}
	sort.Slice(TimeDataBases, func(i, j int) bool {
		return TimeDataBases[i].Id < TimeDataBases[j].Id
	})
	// log.Printf("time = %s", TimeDataBases)
	ids := make([]int, len(TimeDataBases))
	dates := make([]int, len(TimeDataBases))
	hours := make([]int, len(TimeDataBases))
	minutes := make([]int, len(TimeDataBases))
	for i := 0; i < len(TimeDataBases); i++ {
		ids[i] = TimeDataBases[i].Id
		dates[i] = TimeDataBases[i].Date
		hours[i] = TimeDataBases[i].Hour
		minutes[i] = TimeDataBases[i].Minute
	}
	return ids, dates, hours, minutes, nil
}

func SetDateHour(date string, hour string, minute string) error {
	db := Db
	if _, err := db.Exec("UPDATE `time` SET `date` = ?, `hour` = ?, `minute` = ?", date, hour, minute); err != nil {
		return err
	}
	return nil
}

func UpdateDate(id int, date int) error {
	db := Db
	if _, err := db.Exec("UPDATE `time` SET `date` = ? WHERE `id` = ?", date, id); err != nil {
		return err
	}
	return nil
}

func UpdateTime(id int, hour int, minute int) error {
	db := Db
	if _, err := db.Exec("UPDATE `time` SET `hour` = ?, `minute` = ? WHERE `id` = ?", hour, minute, id); err != nil {
		return err
	}
	return nil
}
