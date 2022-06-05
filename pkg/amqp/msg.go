package amqp

type Pattern string

const (
	JudgeTaskNew    = "judge_task:new"
	JudgeTaskUpdate = "judge_task:update"
)

type Message struct {
	Pattern Pattern     `json:"pattern"`
	Data    interface{} `json:"data"`
}
