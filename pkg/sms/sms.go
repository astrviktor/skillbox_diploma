package sms

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func CheckCountry(country string) bool {
	countries := []string{
		"AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH",
		"BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF",
		"BI", "KH", "CM", "CA", "CV", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG", "CD", "CK", "CR",
		"CI", "HR", "CU", "CW", "CY", "CZ", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE", "ET", "FK",
		"FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU",
		"GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN", "ID", "IR", "IQ", "IE", "IM",
		"IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE", "KI", "KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR",
		"LY", "LI", "LT", "LU", "MO", "MK", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX",
		"FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA", "NR", "NP", "NL", "NC", "NZ", "NI", "NE", "NG",
		"NU", "NF", "MP", "NO", "OM", "PK", "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA",
		"RE", "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM", "VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC",
		"SL", "SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS", "SS", "ES", "LK", "SD", "SR", "SJ", "SZ", "SE", "CH",
		"SY", "TW", "TJ", "TZ", "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA", "AE",
		"GB", "US", "UM", "UY", "UZ", "VU", "VE", "VN", "VG", "VI", "WF", "EH", "YE", "ZM", "ZW"}
	return contains(countries, country)
}

func CheckBandwidth(bandwidth string) bool {
	i, err := strconv.Atoi(bandwidth)
	if err != nil {
		return false
	}

	if i >= 0 && i <= 100 {
		return true
	}

	return false
}

func ResponseTime(time string) bool {
	i, err := strconv.Atoi(time)
	if err != nil {
		return false
	}

	if i >= 0 {
		return true
	}

	return false
}

func CheckProvider(provider string) bool {
	providers := []string{"Topolo", "Rond", "Kildy"}
	return contains(providers, provider)
}

func StatusSMS(csvFile string) []SMSData {
	smsData := make([]SMSData, 0)

	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err.Error() + `: ` + csvFile)
		return smsData
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		data := strings.Split(line, ";")

		if len(data) != 4 {
			continue
		}

		if CheckCountry(data[0]) && CheckBandwidth(data[1]) && ResponseTime(data[2]) && CheckProvider(data[3]) {
			elem := SMSData{Country: data[0], Bandwidth: data[1], ResponseTime: data[2], Provider: data[3]}
			smsData = append(smsData, elem)
		}
	}

	return smsData
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
