package resources

import (
	"encoding/json"
	"github.com/faiface/pixel"
	"io/ioutil"
)

type SpritesheetData struct {
	Sprites map[string]SpritesheetSprite
	Picture pixel.Picture
}

type Pivot struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
type Size struct {
	W float64 `json:"w"`
	H float64 `json:"h"`
}

type Rect struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}

type SpritesheetSprite struct {
	FileName         string `json:"filename"`
	Frame            Rect   `json:"frame"`
	Rotated          bool   `json:"rotated"`
	Trimmed          bool   `json:"trimmed"`
	SpriteSourceSize Rect   `json:"spriteSourceSize"`
	SourceSize       Size   `json:"sourceSize"`
	Pivot            Pivot  `json:"pivot"`
}

func (s *SpritesheetData) Init(params SpritesheetParams) error {
	f, err := ioutil.ReadFile(params.DataFile)
	if err != nil {
		return err
	}

	var sprites []SpritesheetSprite
	err = json.Unmarshal(f, &sprites)
	if err != nil {
		return err
	}

	s.Sprites = make(map[string]SpritesheetSprite)
	for _, sprite := range sprites {
		s.Sprites[sprite.FileName] = sprite
	}

	pic, err := loadPicture(params.ImageFile)
	if err != nil {
		return err
	}
	s.Picture = pic
	return nil
}

func (s *SpritesheetData) GetSprite(name string) *pixel.Sprite {
	sData := s.Sprites[name]
	x := sData.Frame.X
	y := s.Picture.Bounds().Max.Y - sData.Frame.Y
	w := x + sData.Frame.W
	h := y - sData.Frame.H
	sprite := pixel.NewSprite(s.Picture, pixel.R(x, y, w, h))
	return sprite
}
