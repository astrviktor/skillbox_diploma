package mms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"skillbox_diploma/pkg/checker"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func CheckMMS(data []MMSData) []MMSData {
	result := make([]MMSData, 0)

	for _, elem := range data {
		if checker.CheckCountry(elem.Country) && checker.CheckBandwidth(elem.Bandwidth) &&
			checker.CheckResponseTime(elem.ResponseTime) && checker.CheckProvider(elem.Provider) {
			result = append(result, elem)
		}
	}

	return result
}

func StatusMMS(url string) []MMSData {
	mmsData := make([]MMSData, 0)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error() + `: ` + url)
		return mmsData
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []MMSData{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return []MMSData{}
	}

	if err := json.Unmarshal(body, &mmsData); err != nil {
		return []MMSData{}
	}

	return CheckMMS(mmsData)
}
