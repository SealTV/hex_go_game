package resources

import (
	"encoding/json"
	"io/ioutil"
)

type Settings struct {
	AppName           string              `json:"name"`
	Version           string              `json:"version"`
	WindowW           float64             `json:"screen_w"`
	WindowH           float64             `json:"screen_h"`
	IsUseVSync        bool                `json:"vsync"`
	LocalizationsPath string              `json:"localization_path"`
	Sprites           string              `json:"sprites_path"`
	Fonts             string              `json:"fonts_path"`
	Spritesheets      []SpritesheetParams `json:"spritesheets"`
}

type SpritesheetParams struct {
	DataFile  string `json:"data_file"`
	ImageFile string `json:"image_file"`
}

func (s *Settings) Init(path string) error {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(f, s)
	return err
}
