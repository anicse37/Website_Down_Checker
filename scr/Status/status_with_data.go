package files

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	mailhandler "github.com/anicse27/Website_Down_Checker/scr/Mail_Handler"
	variables "github.com/anicse27/Website_Down_Checker/scr/Variables"
)

func CheckStatusWithData(URL_Data variables.URL_Data) {
	response, err := http.Get(URL_Data.SiteURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	ResponseCode := response.StatusCode

	data, LastStatusInt, CountInt, err := LastStatus(URL_Data.SiteName)
	if err != nil {
		fmt.Println(err)
		return
	}

	if ResponseCode == LastStatusInt {
		CountInt++
		if CountInt == 2 {
			mailhandler.SendMail(URL_Data, ResponseCode)
		}
		SaveJson(URL_Data.SiteName, data, CountInt, LastStatusInt)
	} else if ResponseCode != LastStatusInt {
		CountInt = 0
		LastStatusInt = ResponseCode
		SaveJson(URL_Data.SiteName, data, CountInt, LastStatusInt)
	} else {
		fmt.Println("Invalid Data")
	}
	fmt.Println(time.Now().Local().Format("01/02/2006 15:04"), "|| Name:", URL_Data.SiteName, "|| Status Code:", ResponseCode, "|| Count:", CountInt)
}
func LastStatus(name string) (JsonData, int, int, error) {
	err := os.MkdirAll(jsonDir, 0755)
	if err != nil {
		return JsonData{}, 0, 0, err
	}

	filename := jsonDir + "/" + name + ".json"

	// Read file
	content, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// Create file with {}
			err = os.WriteFile(filename, []byte(`{}`), 0644)
			if err != nil {
				return JsonData{}, 0, 0, err
			}
			content = []byte(`{}`)
		} else {
			return JsonData{}, 0, 0, err
		}
	}

	data := JsonData{}

	// {} is valid JSON
	if err := json.Unmarshal(content, &data); err != nil {
		return JsonData{}, 0, 0, err
	}

	lastStatus := 0
	count := 0

	if data.Web.LastStatus != "" {
		lastStatus, err = strconv.Atoi(data.Web.LastStatus)
		if err != nil {
			return JsonData{}, 0, 0, err
		}
	}

	if data.Web.Count != "" {
		count, err = strconv.Atoi(data.Web.Count)
		if err != nil {
			return JsonData{}, 0, 0, err
		}
	}

	return data, lastStatus, count, nil
}
