package files

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type JsonDataValues struct {
	LastStatus string `json:"LastStatus"`
	Count      string `json:"Count"`
}
type JsonData struct {
	Web JsonDataValues `json:"Web"`
}

func SaveJson(data JsonData, CountInt, LastStatusInt int) {
	data.Web.Count = strconv.Itoa(CountInt)
	data.Web.LastStatus = strconv.Itoa(LastStatusInt)
	newjsonfile, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	os.Remove("data.json")
	os.WriteFile("data.json", newjsonfile, 0644)
}
