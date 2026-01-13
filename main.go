package main

import (
	"fmt"
	"os"

	"github.com/anicse27/Website_Down_Checker/files"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	files.CheckStatus(os.Getenv("URL"))
}
