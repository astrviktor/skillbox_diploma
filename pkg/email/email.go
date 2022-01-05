package email

import (
	"bufio"
	"fmt"
	"os"
	"skillbox_diploma/pkg/check"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

func ParseEmailData(line string) (EmailData, bool) {
	data := strings.Split(line, ";")

	if len(data) != 3 {
		return EmailData{}, false
	}

	Country := data[0]
	if !check.IsCountry(Country) {
		return EmailData{}, false
	}

	Provider := data[1]
	if !check.IsProviderEmail(Provider) {
		return EmailData{}, false
	}

	DeliveryTime, err := strconv.Atoi(data[2])
	if err != nil {
		return EmailData{}, false
	}

	elem := EmailData{
		Country:      Country,
		Provider:     Provider,
		DeliveryTime: DeliveryTime,
	}

	return elem, true
}

func StatusEmail(csvFile string) []EmailData {
	result := make([]EmailData, 0)

	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err.Error() + `: ` + csvFile)
		return []EmailData{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		elem, ok := ParseEmailData(line)

		if ok {
			result = append(result, elem)
		}
	}

	return result
}
