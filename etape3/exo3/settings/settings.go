package settings

import (
	"encoding/json"
	"io/ioutil"
)

// StrServeurSetting : configurations du serveur
type StrServeurSetting struct {
	Environnement struct {
		Name string `json:"name"`
		HTTP struct {
			ListenPort int    `json:"listenPort"`
			Path       string `json:"path"`
		} `json:"http"`
		ModeDebug   bool `json:"modeDebug"`
		LevelLogger int  `json:"levelLogger"`
	} `json:"environnement"`
}

var settings StrServeurSetting

// LoadSettings : Charge la configuration
func LoadSettings(environnement string) error {

	var err error
	var setup []byte
	var tblSettings []StrServeurSetting

	if setup, err = ioutil.ReadFile("settings.json"); err == nil {
		if err = json.Unmarshal(setup, &tblSettings); err == nil {
			for _, settings = range tblSettings {
				if settings.Environnement.Name == environnement {
					break
				}
			}
		}

	}
	return err
}

// GetSettings : Renvoi la structure charger
func GetSettings() StrServeurSetting {
	return settings
}
