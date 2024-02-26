package sandbox

import "judger/pkg/runner"

type ExecuteOption struct {
	TimeLimit   float64
	MemoryLimit float64
}

type Sandbox interface {
	Exists() bool
	Compile(runner runner.Runner, workSpace string) int
	Execute(runner runner.Runner, workSpace string, test string, option ExecuteOption) int
}
