package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	SimulatorAddr string `yaml:"simulator_addr"`
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
}

var GlobalConfig Config

func NewConfig(file string) Config {
	var config Config

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return GetDefaultConfig()
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return GetDefaultConfig()
	}

	return config
}

func GetDefaultConfig() Config {
	const dir = "/home/astrviktor/golang/src/skillbox_diploma/cmd/data/"
	const addr = "127.0.0.1:9999"

	var config Config

	config.SimulatorAddr = addr
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

	return config
}
