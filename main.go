package main

import (
	"fmt"
	"github.com/SealTV/hex_go_game/resources"
	"github.com/labstack/gommon/log"
)

func main() {
	dirPath := "./data/sprites"
	localizationPath := "./data/localization"
	sMap := make(resources.SpriteMap)
	err := sMap.Init(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	var localizator resources.Localizator
	localizator.Init(localizationPath, resources.En)
	fmt.Println(localizator.GetString("AppTitle"))
	localizator.SetLang(resources.Ru)
	fmt.Println(localizator.GetString("AppTitle"))
	fmt.Println(localizator.GetString("asd"))
}
