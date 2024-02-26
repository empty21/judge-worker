package model

import "judger/pkg/config"

type JudgeTaskResult struct {
	Identifier   string                `json:"identifier"`
	Status       config.TaskStatusEnum `json:"status,omitempty"`
	Tests        []TestResult          `json:"tests,omitempty"`
	ErrorMessage string                `json:"errorMessage,omitempty"`
}

type TestResult struct {
	Identifier string                `json:"identifier"`
	Result     config.TestResultEnum `json:"result"`
	Memory     int64                 `json:"memory"`
	Time       float64               `json:"time"`
}

func NewJudgeTaskStatus(identifier string, status config.TaskStatusEnum, errorMessage string) JudgeTaskResult {
	return JudgeTaskResult{
		Identifier:   identifier,
		Status:       status,
		ErrorMessage: errorMessage,
	}
}

func NewJudgeTaskResult(identifier string, tests []TestResult) JudgeTaskResult {
	return JudgeTaskResult{
		Identifier: identifier,
		Tests:      tests,
	}
}
