package settings

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
func LoadSettings() {
	settings.Environnement.Name = "Dev"
	settings.Environnement.HTTP.ListenPort = 8882
	settings.Environnement.HTTP.Path = "/dev"
	settings.Environnement.ModeDebug = true
	settings.Environnement.LevelLogger = 5
}

// GetSettings : Renvoi la structure charger
func GetSettings() StrServeurSetting {
	return settings
}
