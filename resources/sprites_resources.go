package resources

import (
	"fmt"

	"github.com/faiface/pixel"
	_ "github.com/faiface/pixel/pixelgl"
	_ "golang.org/x/image/colornames"

	"image"
	_ "image/png"
	"io/ioutil"
	"os"
)

type SpriteMap map[string]pixel.Picture

type Sprite struct {
	Name    string        `json:"name"`
	Picture pixel.Picture `json:"picture"`
}

func (smap SpriteMap) Init(path string) error {
	sprites, err := LoadAllSpritesInDirectory(path, "")
	if err != nil {
		return err
	}

	for _, s := range sprites {
		smap[s.Name] = s.Picture
	}

	return nil
}

func (smap SpriteMap) GetSprite(name string) pixel.Picture {
	s, b := smap[name]
	if !b {
		return nil
	}

	return s
}

func LoadAllSpritesInDirectory(path, prefix string) ([]Sprite, error) {
	result := make([]Sprite, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		p := fmt.Sprintf("%s/%s", path, f.Name())
		if f.IsDir() {
			sprites, err := LoadAllSpritesInDirectory(p, fmt.Sprintf("%s/%s", prefix, f.Name()))
			if err != nil {
				return nil, err
			}
			result = append(result, sprites...)
		} else {
			pic, err := loadPicture(p)
			if err != nil {
				return nil, err
			}

			sprite := Sprite{fmt.Sprintf("%s/%s", prefix, f.Name()), pic}
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
