package runner

import (
	"fmt"
	"judger/pkg/config"
	"strings"
)

type _GoRunner struct {
}

func (r *_GoRunner) SandboxImage() string {
	return config.SandboxImageGo
}

func (r *_GoRunner) SourceFileName() string {
	return "main.go"
}

func (r *_GoRunner) ExecutableFileName() string {
	return strings.Replace(r.SourceFileName(), ".go", ".o", -1)
}

func (r *_GoRunner) CompileCommand() string {
	return fmt.Sprintf("go build -o %s %s", r.ExecutableFileName(), r.SourceFileName())
}

func (r *_GoRunner) ExecuteCommand() string {
	return fmt.Sprintf("./%s", r.ExecutableFileName())
}

func init() {
	registry[config.LangGo] = &_GoRunner{}
}
