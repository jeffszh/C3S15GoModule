package model

import (
	"gopkg.in/yaml.v2"
	"os"
)

const (
	configFileName = "AppConf.yaml"
)

type AppConfigStruct struct {
	SoldierText string `yaml:"soldier_text" json:"soldier_text,omitempty"`
	CannonText  string `yaml:"cannon_text" json:"cannon_text,omitempty"`
	AiDepth     int    `yaml:"ai_depth" json:"ai_depth,omitempty"`
}

var AppConfig AppConfigStruct

func init() {
	bytes, err := os.ReadFile(configFileName)
	if err != nil {
		bytes = createDefaultConfigFile()
	}
	_ = yaml.Unmarshal(bytes, &AppConfig)
}

func createDefaultConfigFile() []byte {
	conf := AppConfigStruct{
		SoldierText: "兵",
		CannonText:  "炮",
		AiDepth:     5,
	}
	result, _ := yaml.Marshal(conf)
	_ = os.WriteFile(configFileName, result, 0666)
	return result
}
