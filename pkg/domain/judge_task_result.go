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
	TestResultIE                 = "IE"
)

type JudgeTaskResult struct {
	SubmissionId int64          `json:"submissionId"`
	Status       TaskStatusEnum `json:"status,omitempty"`
	Tests        []TestResult   `json:"tests,omitempty"`
}

type TestResult struct {
	TestUuid string         `json:"testUuid"`
	Result   TestResultEnum `json:"result"`
	Memory   int64          `json:"memory"`
	Time     float64        `json:"time"`
}

func NewJudgeTaskStatus(submissionId int64, status TaskStatusEnum) JudgeTaskResult {
	return JudgeTaskResult{
		SubmissionId: submissionId,
		Status:       status,
	}
}

func NewJudgeTaskResult(submissionId int64, tests []TestResult) JudgeTaskResult {
	return JudgeTaskResult{
		SubmissionId: submissionId,
		Tests:        tests,
	}
}
