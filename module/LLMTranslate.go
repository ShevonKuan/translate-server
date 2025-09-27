package module

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/abadojack/whatlanggo"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func LLMTranslate(i *InputObj, c *gin.Context) (*OutputObj, int, error) {

	sourceLang := i.SourceLang
	targetLang := i.TargetLang
	translateText := i.SourceText
	API := i.API
	token := i.Token
	if token == "" {
		log.Println("Token not provided for LLM translation")
		return nil, 500, errors.New("Token is required for LLM translation")

	}
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
		switch API {
		case "deepl":
			log.Println("Using DeepL API for LLM translation")
			return nil, 500, errors.New("DeepL API for LLM translation is not supported yet")
		default: // GLM
			u := "https://open.bigmodel.cn/api/paas/v4/chat/completions"
			method := "POST"
			headers := map[string]string{
				"Authorization": "Bearer " + token,
				"Content-Type":  "application/json",
			}
			reqData, _ := json.Marshal(NewGLMRequest(*i))

			client := &http.Client{}
			req, err := http.NewRequest(method, u, bytes.NewBuffer(reqData))
			for k, v := range headers {
				req.Header.Set(k, v)
			}

			if err != nil {
				fmt.Println(err)
				return nil, 0, err
			}
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
				return nil, 0, err
			}
			if resp.StatusCode != 200 {
				body, _ := io.ReadAll(resp.Body)
				log.Println("LLM API error:", resp.StatusCode, string(body))
				return nil, resp.StatusCode, errors.New("LLM API error: " + string(body))
			}
			defer resp.Body.Close()
			// check stream if true
			if i.Stream == "true" { // handle stream response
				// --- 开始流式响应 ---
				// 设置响应头，告知浏览器这是一个 SSE (Server-Sent Events) 流
				c.Header("Content-Type", "text/event-stream; charset=utf-8")
				c.Header("Cache-Control", "no-cache")
				c.Header("Connection", "keep-alive")
				c.Header("Transfer-Encoding", "chunked")
				// 如果前端是不同域，需要设置 CORS 头
				c.Header("Access-Control-Allow-Origin", "*")

				// 使用 bufio.Scanner 逐行读取 Ollama 的响应流
				scanner := bufio.NewScanner(resp.Body)
				color.Cyan("Translate Begin")
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" {
						continue // 跳过空行
					}

					// 解析返回的 JSON 行
					res := gjson.Parse(line[6:]) // 去掉 "data: " 前缀
					result := res.Get("choices.0.delta.content").String()
					if res.Get("choices.0.finish_reason").String() == "stop" {
						c.SSEvent("data", res.Get("usage").String())
						c.Writer.Flush()
						c.SSEvent("data", "[DONE]")
						c.Writer.Flush()
						break
					}
					fmt.Print(result)
					// 将返回的文本块通过 SSE 发送给前端

					// 使用 c.SSEvent 发送数据事件
					// 注意：gin 的 SSE 实现会自动添加 "data: " 前缀和换行符
					c.SSEvent("data", result)
					// 强制刷新响应缓冲区，立即将数据发送给客户端
					c.Writer.Flush()

				}
				color.Cyan("\nTranslate Finished ")
				return nil, http.StatusOK, nil
			} else {

				body, _ := io.ReadAll(resp.Body)
				res := gjson.ParseBytes(body)
				result := res.Get("choices.0.message.content").String()
				color.Cyan("Translate: ", result, res.Get("usage").String())

				c.JSON(http.StatusOK, gin.H{
					"code":   http.StatusOK,
					"data":   result,
					"engine": "GLM",
					"usage":  res.Get("usage").String(),
				})
				return nil, http.StatusOK, nil
			}
		}
	}
}
