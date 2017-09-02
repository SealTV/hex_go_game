package resources

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"io/ioutil"
)

type LangType string

const (
	En LangType = "en"
	Ru LangType = "ru"
)

type Lang struct {
	Name     string `json:"lang"`
	FileName string `json:"path"`
}

type LocalizationSting struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Localizator struct {
	Stings      map[LangType]map[string]string
	CurrentLang LangType
}

func (l *Localizator) Init(path string, defaultLang LangType) {
	l.CurrentLang = defaultLang
	languages, err := loadLanguagesSettings(path + "/langs.json")
	if err != nil {
		log.Fatal(err)
	}

	l.Stings = make(map[LangType]map[string]string)
	for _, lang := range languages {
		l.Stings[LangType(lang.Name)] = make(map[string]string)
		strings, err := loadLangStrings(path + "/" + lang.FileName)
		if err != nil {
			log.Fatal(err)
		}

		for _, s := range strings {
			l.Stings[LangType(lang.Name)][s.Key] = s.Value
		}
	}
}

func (l *Localizator) SetLang(lang LangType) {
	l.CurrentLang = lang
}

func (l *Localizator) GetString(key string) string {
	if l.Stings[l.CurrentLang] != nil {
		if l.Stings[l.CurrentLang][key] != "" {
			return l.Stings[l.CurrentLang][key]
		}
	}
	return ""
}

func loadLanguagesSettings(path string) ([]Lang, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var languages []Lang
	err = json.Unmarshal(file, &languages)
	if err != nil {
		return nil, err
	}

	return languages, nil
}

func loadLangStrings(path string) ([]LocalizationSting, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var stings []LocalizationSting
	err = json.Unmarshal(file, &stings)
	if err != nil {
		return nil, err
	}

	return stings, nil
}
