package sandbox

import (
	"judger/pkg/domain"
	"judger/pkg/runner"
)

type Sandbox interface {
	Compile(runner runner.Runner, workDir string) (error, int)
	Execute(runner runner.Runner, workDir string, test domain.Test, limitation domain.TaskLimitation) (error, *domain.TestResult)
}
