package module

import (
	"fmt"
	"strconv"
)

type LLMPrompt struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type LLMMessages []LLMPrompt
type ThinkingSwitch struct {
	Type string `json:"type"`
}

// GLMRequest GLM 智谱大模型

// Available Model: glm-4.5, glm-4.5-air, glm-4.5-x, glm-4.5-airx, glm-4.5-flash, glm-4-plus, glm-4-air-250414, glm-4-airx, glm-4-flashx, glm-4-flashx-250414, glm-z1-air, glm-z1-airx, glm-z1-flash, glm-z1-flashx
type GLMRequest struct {
	ModelName   string         `json:"model"`
	Temperature string         `json:"temperature,omitempty"`
	MaxTokens   int64          `json:"max_tokens,omitempty"`
	Messages    LLMMessages    `json:"messages"`
	Stream      bool           `json:"stream,omitempty"`
	Thinking    ThinkingSwitch `json:"thinking,omitempty"`
}

func NewGLMRequest(obj InputObj) *GLMRequest {
	// change inputObj to GLMRequest
	ModelName := obj.ModelName
	if ModelName == "" {
		ModelName = "glm-4.5-flash"

	}
	Temperature := obj.Temperature
	if Temperature == "" {
		Temperature = "0.1"
	}
	MaxTokens := obj.MaxTokens
	if MaxTokens == "" {
		MaxTokens = "2048"
	}
	Message := LLMMessages{}
	if obj.Prompt == "" {
		c1 := fmt.Sprintf("As an academic expert with specialized knowledge in various fields, please provide a proficient and precise translation to %s .\n\n## Translation Rules\n1. Output only the translated content, without explanations or additional content (such as \"Here's the translation:\" or \"Translation as follows:\")\n2. The returned translation must maintain exactly the same number of paragraphs and format as the original text\n3. If the text contains math symbols, try rewrite them into LaTex formula\n4. For content that should not be translated (such as proper nouns, code, etc.), keep the original text.\n", obj.TargetLang)
		c2 := fmt.Sprintf("Please translate the following content from %s to %s:\n\n %s", obj.SourceLang, obj.TargetLang, obj.SourceText)
		Message = LLMMessages{
			{Role: "system", Content: c1},
			{Role: "user", Content: c2},
		}
	} else {
		Message = LLMMessages{
			{Role: "system", Content: obj.Prompt},
			{Role: "user", Content: obj.SourceText},
		}
	}
	m, _ := strconv.ParseInt(MaxTokens, 10, 32)

	return &GLMRequest{
		ModelName:   ModelName,
		Temperature: Temperature,
		MaxTokens:   m,
		Messages:    Message,
		Stream:      obj.Stream == "true",
		Thinking:    ThinkingSwitch{Type: "disabled"},
	}
}
