package module

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/abadojack/whatlanggo"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

func GoogleTranslate(i *InputObj) (*OutputObj, int, error) {

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
		u := "https://translate.googleapis.com/translate_a/single"
		u += "?client=gtx&sl=" + sourceLang + "&tl=" + targetLang + "&dt=t&q=" + url.QueryEscape(translateText)
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, u, nil)

		if err != nil {
			fmt.Println(err)
			return nil, 0, err
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return nil, 0, err
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		res := gjson.ParseBytes(body)
		// display response
		// fmt.Println(res)
		// print translate text
		var output = ""
		res.Get("0").ForEach(func(key, value gjson.Result) bool {
			output += value.Get("0").String()
			return true
		})

		color.Cyan("Translate: " + output)
		if res.Get("error.code").String() == "-32600" {
			log.Println(res.Get("error").String())
			return nil, resp.StatusCode, errors.New("invalid targetLang")
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			log.Println("Too many requests")
			return nil, resp.StatusCode, errors.New("too many requests")
		}
		var alternatives []string
		res.Get("result.texts.0.alternatives").ForEach(func(key, value gjson.Result) bool {
			alternatives = append(alternatives, value.Get("text").String())
			return true
		})

		return &OutputObj{
			TransText:    output,
			Alternatives: alternatives,
		}, http.StatusOK, nil

	}
}
