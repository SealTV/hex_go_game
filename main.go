package main

import (
	"fmt"
	"github.com/SealTV/hex_go_game/game"
	"github.com/SealTV/hex_go_game/resources"
	"github.com/faiface/pixel/pixelgl"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {
	fmt.Println("Try read settings")
	if len(os.Args) <= 1 {
		panic("Settings path is empty")
	}

	fmt.Println("Try read settings")
	var settings resources.Settings
	err := settings.Init(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Try load localization strings")
	var localization resources.Localizator
	localization.Init(settings.LocalizationsPath, resources.En)

	fmt.Println("Try load spritesheets")
	var sh resources.SpritesheetData
	err = sh.Init(settings.Spritesheets[0])
	if err != nil {
		panic(err)
	}

	fmt.Println("Try load sprites")
	sprites := make(resources.SpriteMap)
	err = sprites.Init(settings.Sprites)
	if err != nil {
		panic(err)
	}

	fmt.Println("Try start application")
	var game = game.Game{settings, localization, sprites, sh}
	pixelgl.Run(game.Run)
}
