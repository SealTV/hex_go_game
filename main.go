package main

import (
	"fmt"
	"github.com/SealTV/hex_go_game/resources"
	"github.com/labstack/gommon/log"
)

func main() {
	dirPath := "./data/sprites"
	sprites, err := resources.LoadAllSpritesInDirrectory(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range sprites {
		fmt.Println(s.Name)
	}

}
