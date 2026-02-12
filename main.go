package main

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	databasehandler "github.com/anicse27/Website_Down_Checker/scr/Database_Handler"
	files "github.com/anicse27/Website_Down_Checker/scr/Status"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
}

var (
	DSN = "root:Gtech@17@tcp(127.0.0.1:3306)/Website_Down_Checker?charset=utf8mb4&parseTime=true&loc=Local"
)

func main() {
	DB, err := databasehandler.StartDatabase(DSN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer DB.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	URL, err := databasehandler.GetAllURLs(DB, ctx)
	if err != nil {
		fmt.Println(err)
	}
	for _, url := range URL {
		files.CheckStatusWithData(url)
	}
}
