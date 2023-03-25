package domain

type TestResultEnum string
type TaskStatusEnum string

const (
	TaskStatusIP TaskStatusEnum = "IP"
	TaskStatusCE                = "CE"
	TaskStatusIE                = "IE"
)

const (
	TestResultAC  TestResultEnum = "AC"
	TestResultWA                 = "WA"
	TestResultRTE                = "RTE"
	TestResultTLE                = "TLE"
)

type JudgeTaskResult struct {
	Uid    string         `json:"uid"`
	Status TaskStatusEnum `json:"status,omitempty"`
	Tests  []TestResult   `json:"tests,omitempty"`
}

type TestResult struct {
	TestUuid string         `json:"testUuid"`
	Result   TestResultEnum `json:"result"`
	Memory   int64          `json:"memory"`
	Time     float64        `json:"time"`
}

func NewJudgeTaskStatus(uid string, status TaskStatusEnum) JudgeTaskResult {
	return JudgeTaskResult{
		Uid:    uid,
		Status: status,
	}
}

func NewJudgeTaskResult(uid string, tests []TestResult) JudgeTaskResult {
	return JudgeTaskResult{
		Uid:   uid,
		Tests: tests,
	}
}
