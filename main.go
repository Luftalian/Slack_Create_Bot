package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Luftalian/Slack_Create_Bot/cronAction"
	"github.com/Luftalian/Slack_Create_Bot/databaseConnect"
	"github.com/Luftalian/Slack_Create_Bot/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("*/30 * * * *", cronAction.CronActionFunc)
	c.Start()
	log.Printf("cron start")

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
		// return err
	}
	databaseConnect.Db = db
	log.Printf("Connected to Database")

	// err := databaseConnect.DatabaseConnectFunc()
	// if err != nil {
	// 	log.Fatalf("Cannot Connect to Database: %s", err)
	// }

	// cronAction.CronActionFuncNoTime()

	e := echo.New()
	// e.POST("/", router.HandleRequestAndReturnJSON) // Check Challenge
	e.POST("/", router.HandlePostMessageJSON) // EVENT API
	// e.POST("/api/command", router.HandleCommand)

	e.Logger.Fatal(e.Start(":8080"))
}
