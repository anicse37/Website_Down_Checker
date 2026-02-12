package files

import (
	"encoding/json"
	"os"
	"strconv"
)

const jsonDir = "json_files"

type JsonDataValues struct {
	LastStatus string `json:"LastStatus"`
	Count      string `json:"Count"`
}
type JsonData struct {
	Web JsonDataValues `json:"Web"`
}

func SaveJson(name string, data JsonData, count int, lastStatus int) error {
	err := os.MkdirAll(jsonDir, 0755)
	if err != nil {
		return err
	}

	data.Web.Count = strconv.Itoa(count)
	data.Web.LastStatus = strconv.Itoa(lastStatus)

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	filename := jsonDir + "/" + name + ".json"
	return os.WriteFile(filename, bytes, 0644)
}
