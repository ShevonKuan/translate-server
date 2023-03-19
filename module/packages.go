package module

type InputObj struct {
	SourceText string `json:"text"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
}

type OutputObj struct {
	TransText    string `json:"text"`
	Alternatives []string
}
