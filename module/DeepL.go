package module

import (
	"errors"
	"log"
	"net/http"
	"strings"

	translate "github.com/OwO-Network/DeepLX/translate"
	"github.com/abadojack/whatlanggo"
	"github.com/fatih/color"
)

func DeepLTranslate(i *InputObj) (*OutputObj, int, error) {

	sourceLang := i.SourceLang
	targetLang := i.TargetLang
	translateText := i.SourceText

	// print lang
	color.Red("%s -> %s", sourceLang, targetLang)
	// print source text
	color.Green("Source: " + translateText)

	if sourceLang == "" {
		lang := whatlanggo.DetectLang(translateText)
		deepLLang := strings.ToUpper(lang.Iso6391())
		sourceLang = deepLLang
	}
	if targetLang == "" {
		targetLang = "EN"
	}
	if translateText == "" {
		return nil, 0, errors.New("empty text")
	} else {
		result, err := translate.TranslateByDeepLX(
			sourceLang, targetLang, translateText, "", "", "")

		color.Cyan("Translate: ", result.Data)
		if err != nil || result.Data == "" {
			log.Println(err)
			return nil, 500, errors.New("Translate API error: " + result.Message)
		}

		return &OutputObj{
			TransText: result.Data,
		}, http.StatusOK, nil

	}
}
