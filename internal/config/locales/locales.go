package locales

import (
	"github.com/cubiest/jibberjabber"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var (
	default_language = language.English
)

type Locales struct {
	Localizer *i18n.Localizer
}

func NewLocales() Locales {
	return Locales{
		Localizer: getLocalizer(),
	}
}

func getLocalizer() *i18n.Localizer {
	bundle := i18n.NewBundle(default_language)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.MustLoadMessageFile("./internal/config/locales/en.yaml")
	bundle.MustLoadMessageFile("./internal/config/locales/fr.yaml")
	return i18n.NewLocalizer(bundle, getLanguage().String(), default_language.String())
}

func getLanguage() language.Tag {
	userLanguage, err := jibberjabber.DetectLanguage()
	if err != nil {
		return default_language
	}
	return language.Make(userLanguage)
}

func (locales Locales) L(msgId string) string {
	return locales.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: msgId})
}
