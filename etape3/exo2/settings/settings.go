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
func LoadSettings() error {

	var err error
	var setup []byte

	if setup, err = ioutil.ReadFile("settings.json"); err == nil {
		err = json.Unmarshal(setup, &settings)
	}
	return err
}

// GetSettings : Renvoi la structure charger
func GetSettings() StrServeurSetting {
	return settings
}
