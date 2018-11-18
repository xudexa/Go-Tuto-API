package main

import (
	"defis_Exo/etape3/exo3/settings"
	"fmt"
	"log"
)

func main() {
	var appSetting settings.StrServeurSetting
	var err error
	if err = settings.LoadSettings("Prod"); err == nil {
		appSetting = settings.GetSettings()
		fmt.Println(appSetting)
	} else {
		log.Fatal(err.Error())
	}

}
