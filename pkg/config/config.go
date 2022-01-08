package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Addr          string `yaml:"addr"`
	SMSFile       string `yaml:"sms_file"`
	MMSAddr       string `yaml:"mms_addr"`
	MMSFile       string `yaml:"mms_file"`
	VoiceCallFile string `yaml:"voice_call_file"`
	EmailFile     string `yaml:"email_file"`
	BillingFile   string `yaml:"billing_file"`
	SupportAddr   string `yaml:"support_addr"`
	SupportFile   string `yaml:"support_file"`
	IncidentAddr  string `yaml:"incident_addr"`
	IncidentFile  string `yaml:"incident_file"`
	WebDir        string `yaml:"web_dir"`
}

var GlobalConfig Config

func NewConfig(file string) Config {
	var config Config

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return GetDefaultConfig()
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err.Error())
		return GetDefaultConfig()
	}

	return config
}

func GetDefaultConfig() Config {
	fmt.Println("get default config")

	const dir = "/home/astrviktor/golang/src/skillbox_diploma/cmd/data/"
	const addr = ":9999"

	var config Config

	config.Addr = addr
	config.SMSFile = dir + "sms.data"
	config.MMSAddr = "http://" + addr + "/mms"
	config.MMSFile = dir + "mms.json"
	config.VoiceCallFile = dir + "voice.data"
	config.EmailFile = dir + "email.data"
	config.BillingFile = dir + "billing.data"
	config.SupportAddr = "http://" + addr + "/support"
	config.SupportFile = dir + "support.json"
	config.IncidentAddr = "http://" + addr + "/incident"
	config.IncidentFile = dir + "incident.json"
	config.WebDir = "/home/astrviktor/golang/src/skillbox_diploma/web"

	return config
}
