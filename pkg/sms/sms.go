package sms

import (
	"bufio"
	"fmt"
	"os"
	"skillbox_diploma/pkg/checker"
	"strings"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func StatusSMS(csvFile string) []SMSData {
	smsData := make([]SMSData, 0)

	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err.Error() + `: ` + csvFile)
		return []SMSData{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		data := strings.Split(line, ";")

		if len(data) != 4 {
			continue
		}

		if checker.CheckCountry(data[0]) && checker.CheckBandwidth(data[1]) && checker.CheckResponseTime(data[2]) &&
			checker.CheckProvider(data[3]) {
			elem := SMSData{Country: data[0], Bandwidth: data[1], ResponseTime: data[2], Provider: data[3]}
			smsData = append(smsData, elem)
		}
	}

	return smsData
}
