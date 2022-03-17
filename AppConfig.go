package model

import (
	"gopkg.in/yaml.v2"
	"os"
)

const (
	configFileName = "AppConf.yaml"
)

type AppConfigStruct struct {
	AppTitle            string     `yaml:"app_title" json:"app_title"`
	SoldierText         string     `yaml:"soldier_text" json:"soldier_text,omitempty"`
	CannonText          string     `yaml:"cannon_text" json:"cannon_text,omitempty"`
	SoldierPlayType     PlayerType `yaml:"soldier_play_type" json:"soldier_play_type"`
	CannonPlayType      PlayerType `yaml:"cannon_play_type" json:"cannon_play_type"`
	AiDepth             int        `yaml:"ai_depth" json:"ai_depth,omitempty"`
	PlayerTypeHumanText string     `yaml:"player_type_human_text" json:"player_type_human_text"`
	PlayerTypeAIText    string     `yaml:"player_type_ai_text" json:"player_type_ai_text"`
	PlayerTypeNetText   string     `yaml:"player_type_net_text" json:"player_type_net_text"`
}

var AppConfig AppConfigStruct

func init() {
	AppConfig = createDefaultConfig()
	bytes, err := os.ReadFile(configFileName)
	if err == nil {
		_ = yaml.Unmarshal(bytes, &AppConfig)
	}
	saveConfigFile()
}

func createDefaultConfig() AppConfigStruct {
	return AppConfigStruct{
		AppTitle:            "三炮十五兵 Go语言版",
		SoldierText:         "兵",
		CannonText:          "炮",
		PlayerTypeHumanText: "人脑",
		PlayerTypeAIText:    "電腦",
		PlayerTypeNetText:   "网友",
		SoldierPlayType:     1,
		CannonPlayType:      0,
		AiDepth:             5,
	}
}

func saveConfigFile() {
	bytes, _ := yaml.Marshal(AppConfig)
	_ = os.WriteFile(configFileName, bytes, 0666)
}
