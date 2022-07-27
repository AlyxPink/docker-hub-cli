package locales

import (
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

type Locales struct {
	Localizer *i18n.Localizer
}

func GetLocalizer() *i18n.Localizer {
	default_language := language.English
	bundle := i18n.NewBundle(default_language)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.MustLoadMessageFile("./internal/config/locales/en.yaml")
	bundle.MustLoadMessageFile("./internal/config/locales/fr.yaml")
	return i18n.NewLocalizer(bundle, os.Getenv("LANG"), default_language.String())
}

func (locales Locales) Localize(msgId string) string {
	return locales.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: msgId})
}
