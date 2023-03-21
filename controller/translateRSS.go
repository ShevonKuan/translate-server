package controller

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/ShevonKuan/translate-server/module"
	"github.com/abadojack/whatlanggo"
	"github.com/beevik/etree"
	"github.com/gin-gonic/gin"
)

type Rss struct {
	XMLName xml.Name `xml:"RDF"`
	Channel Channel  `xml:"channel"`
}
type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Desc    string   `xml:"description"`
	Items   []Item   `xml:"item"`
}
type Item struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Desc    string   `xml:"description"`
}

// 翻译函数
func translate(i string, engine string) string {
	lang := whatlanggo.DetectLang(i)
	sourceLang := strings.ToUpper(lang.Iso6391())
	// 调用翻译函数
	input := &module.InputObj{
		SourceText: strings.ReplaceAll(i, "\n", " "),
		SourceLang: sourceLang,
		TargetLang: "zh",
	}
	output, _, err := module.Engine[engine](input)
	// 返回翻译结果
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return output.TransText
}

// 获取xml内容并解析为结构体
func getRss(url string) (*etree.Document, error) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	rss := etree.NewDocument()
	if err := rss.ReadFromBytes(body); err != nil {
		return nil, err
	}
	return rss, nil
}
func RSStranslate(url string, engine string) (*etree.Document, error) {
	// 获取xml内容并解析为结构体
	rss, err := getRss(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 翻译 channel
	title := rss.FindElement("//channel/title")
	title.SetText(translate(title.Text(), engine) + title.Text())
	description := rss.FindElement("//channel/description")
	description.SetText(translate(description.Text(), engine) + description.Text())
	// 翻译 item
	var wg sync.WaitGroup
	items := rss.FindElements("//item")

	for _, item := range items {
		wg.Add(1)
		// 并发翻译
		go func(item *etree.Element) {
			defer wg.Done()
			title := item.FindElement("title")
			title.SetText(translate(title.Text(), engine) + title.Text())
			description := item.FindElement("description")
			fmt.Println(description.Text())
			description.SetText(translate(description.Text(), engine) + description.Text())
		}(item)
	}
	wg.Wait()
	return rss, nil
}
func TranslateRSS(c *gin.Context) {
	translateEngine, err := module.GetEngine(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err,
		})
	}

	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "url not found",
		})
		return
	}
	resp, err := RSStranslate(url, translateEngine)
	if err != nil {
		c.XML(http.StatusNotFound, gin.H{
			"code":    http.StatusServiceUnavailable,
			"message": err,
		})
	}
	output, _ := resp.WriteToString()
	c.Writer.Write([]byte(output))
}
