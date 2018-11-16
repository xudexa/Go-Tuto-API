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
