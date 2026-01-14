package files

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func CheckStatus(URL string) {
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	ResponseCode := response.StatusCode

	data, LastStatusInt, CountInt, err := LastStatus()
	if err != nil {
		fmt.Println(err)
		return
	}

	if ResponseCode == LastStatusInt {
		CountInt++
		if CountInt == 2 {
			SendMail(ResponseCode)
		}
		SaveJson(data, CountInt, LastStatusInt)
	} else if ResponseCode != LastStatusInt {
		CountInt = 0
		LastStatusInt = ResponseCode
		SaveJson(data, CountInt, LastStatusInt)
	} else {
		fmt.Println("Invalid Data")
	}
	fmt.Println(time.Now().Local().Format("01/02/2006 15:04"), "|| Status Code:", ResponseCode, "|| Count:", CountInt)
}

func LastStatus() (JsonData, int, int, error) {
	LastStatus, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return JsonData{}, 0, 0, err
	}
	data := JsonData{}
	err = json.Unmarshal(LastStatus, &data)
	if err != nil {
		fmt.Println(err)
		return JsonData{}, 0, 0, err

	}

	LastStatusInt, err := strconv.Atoi(data.Web.LastStatus)
	if err != nil {
		fmt.Println(err)
		return JsonData{}, 0, 0, err

	}
	CountInt, err := strconv.Atoi(data.Web.Count)
	if err != nil {
		fmt.Println(err)
		return JsonData{}, 0, 0, err

	}
	return data, LastStatusInt, CountInt, nil
}
