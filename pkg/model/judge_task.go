package model

type TaskLimitation struct {
	TimeLimit   float64 `json:"timeLimit"`
	MemoryLimit float64 `json:"memoryLimit"`
}

type JudgeTask struct {
	Identifier   string `json:"identifier"`
	Source       string `json:"source"`
	LanguageCode string `json:"languageCode"`
	Tests        []Test `json:"tests"`
	TaskLimitation
}

type Test struct {
	Identifier string `json:"identifier"`
	InputUri   string `json:"inputUri"`
	OutputUri  string `json:"outputUri"`
	TaskLimitation
}
