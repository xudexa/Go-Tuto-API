package main

import (
	"defis_Exo/etape3/exo1/settings"
	"fmt"
)

func main() {
	var appSetting settings.StrServeurSetting
	settings.LoadSettings()
	appSetting = settings.GetSettings()

	fmt.Println(appSetting)
}
