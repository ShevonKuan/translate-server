package module

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type InputObj struct {
	SourceText string `json:"text"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
}

type OutputObj struct {
	TransText    string `json:"text"`
	Alternatives []string
}

type TranslateEngine func(i *InputObj) (*OutputObj, int, error)

var (
	Engine = map[string]TranslateEngine{
		"deepl":  DeepLTranslate,
		"google": GoogleTranslate,
	}
)

// deepl struct
type deeplLang struct {
	SourceLangUserSelected string `json:"source_lang_user_selected"`
	TargetLang             string `json:"target_lang"`
}

type deeplCommonJobParams struct {
	WasSpoken    bool   `json:"wasSpoken"`
	TranscribeAS string `json:"transcribe_as"`
	// RegionalVariant string `json:"regionalVariant"`
}

type deeplParams struct {
	Texts           []deeplText          `json:"texts"`
	Splitting       string               `json:"splitting"`
	Lang            deeplLang            `json:"lang"`
	Timestamp       int64                `json:"timestamp"`
	CommonJobParams deeplCommonJobParams `json:"commonJobParams"`
}

type deeplText struct {
	Text                string `json:"text"`
	RequestAlternatives int    `json:"requestAlternatives"`
}

type deeplPostData struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	ID      int64       `json:"id"`
	Params  deeplParams `json:"params"`
}

func deeplInitData(sourceLang string, targetLang string) *deeplPostData {
	return &deeplPostData{
		Jsonrpc: "2.0",
		Method:  "LMT_handle_texts",
		Params: deeplParams{
			Splitting: "newlines",
			Lang: deeplLang{
				SourceLangUserSelected: sourceLang,
				TargetLang:             targetLang,
			},
			CommonJobParams: deeplCommonJobParams{
				WasSpoken:    false,
				TranscribeAS: "",
				// RegionalVariant: "en-US",
			},
		},
	}
}

func getICount(translateText string) int64 {
	return int64(strings.Count(translateText, "i"))
}

func GetRandomNumber() int64 {
	rand.Seed(time.Now().Unix())
	num := rand.Int63n(99999) + 8300000
	return num * 1000
}

func getTimeStamp(iCount int64) int64 {
	ts := time.Now().UnixMilli()
	if iCount != 0 {
		iCount = iCount + 1
		return ts - ts%iCount + iCount
	} else {
		return ts
	}
}

// query params processing

func GetEngine(q interface{}) (string, error) {
	// class classify
	var translateEngine string
	c, ok := q.(*gin.Context)
	if ok {
		translateEngine = c.Query("engine")
	} else {
		r, ok := q.(*http.Request)
		if ok {
			translateEngine = r.URL.Query().Get("engine")
		} else {
			return "google", nil
		}
	}

	for e := range Engine {
		if e == translateEngine {
			return translateEngine, nil
		}
	}

	return "google", nil
}
