package module

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/abadojack/whatlanggo"
	"github.com/andybalholm/brotli"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

func DeepLTranslate(i *InputObj) (*OutputObj, int, error) {

	// create a random id
	id := GetRandomNumber()
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
		url := "https://www2.deepl.com/jsonrpc"
		id = id + 1
		postData := deeplInitData(sourceLang, targetLang)
		text := deeplText{
			Text:                translateText,
			RequestAlternatives: 3,
		}
		// set id
		postData.ID = id
		// set text
		postData.Params.Texts = append(postData.Params.Texts, text)
		// set timestamp
		postData.Params.Timestamp = getTimeStamp(getICount(translateText))
		post_byte, _ := json.Marshal(postData)
		postStr := string(post_byte)

		// add space if necessary
		if (id+5)%29 == 0 || (id+3)%13 == 0 {
			postStr = strings.Replace(postStr, "\"method\":\"", "\"method\" : \"", -1)
		} else {
			postStr = strings.Replace(postStr, "\"method\":\"", "\"method\": \"", -1)
		}

		post_byte = []byte(postStr)
		reader := bytes.NewReader(post_byte)
		request, err := http.NewRequest("POST", url, reader)
		if err != nil {
			log.Println(err)
			return nil, 0, err
		}

		// Set Headers
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept", "*/*")
		request.Header.Set("x-app-os-name", "iOS")
		request.Header.Set("x-app-os-version", "16.3.0")
		request.Header.Set("Accept-Language", "en-US,en;q=0.9")
		request.Header.Set("Accept-Encoding", "gzip, deflate, br")
		request.Header.Set("x-app-device", "iPhone13,2")
		request.Header.Set("User-Agent", "DeepL-iOS/2.9.1 iOS 16.3.0 (iPhone13,2)")
		request.Header.Set("x-app-build", "510265")
		request.Header.Set("x-app-version", "2.9.1")
		request.Header.Set("Connection", "keep-alive")

		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			log.Println(err)
			return nil, 0, err
		}
		defer resp.Body.Close()

		var bodyReader io.Reader
		switch resp.Header.Get("Content-Encoding") {
		case "br":
			bodyReader = brotli.NewReader(resp.Body)
		default:
			bodyReader = resp.Body
		}

		body, err := io.ReadAll(bodyReader)
		res := gjson.ParseBytes(body)
		// display response
		// fmt.Println(res)
		// print translate text
		color.Cyan("Translate: " + res.Get("result.texts.0.text").String())
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
			TransText:    res.Get("result.texts.0.text").String(),
			Alternatives: alternatives,
		}, http.StatusOK, nil

	}
}
