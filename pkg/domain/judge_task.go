package domain

type TaskLimitation struct {
	TimeLimit   float64 `json:"timeLimit"`
	MemoryLimit float64 `json:"memoryLimit"`
}

type JudgeTask struct {
	Uid          string `json:"uid"`
	Source       string `json:"source"`
	LanguageCode string `json:"languageCode"`
	TaskLimitation
	Tests []Test `json:"tests"`
}

type Test struct {
	Uuid      string `json:"uuid"`
	InputUri  string `json:"inputUri"`
	OutputUri string `json:"outputUri"`
}
