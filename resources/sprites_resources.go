package resources

import (
	"fmt"
	"github.com/faiface/pixel"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"
)

type Sprite struct {
	Name    string        `json:"name"`
	Picture pixel.Picture `json:"picture"`
}

func LoadAllSpritesInDirrectory(path string) ([]Sprite, error) {
	result := make([]Sprite, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		p := fmt.Sprintf("%s/%s", path, f.Name())
		if f.IsDir() {
			sprites, err := LoadAllSpritesInDirrectory(p)
			if err != nil {
				return nil, err
			}
			result = append(result, sprites...)
		} else {
			pic, err := loadPicture(p)
			if err != nil {
				return nil, err
			}

			sprite := Sprite{f.Name(), pic}
			result = append(result, sprite)
		}
	}
	return result, nil
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}
