package runner

import (
	"fmt"
	"judger/pkg/config"
	"strings"
)

type _CRunner struct {
}

func (r *_CRunner) SandboxImage() string {
	return config.SandboxImageC
}

func (r *_CRunner) SourceFileName() string {
	return "main.c"
}

func (r *_CRunner) ExecutableFileName() string {
	return strings.Replace(r.SourceFileName(), ".c", ".o", -1)
}

func (r *_CRunner) CompileCommand() string {
	return fmt.Sprintf("gcc -o %s %s", r.ExecutableFileName(), r.SourceFileName())
}

func (r *_CRunner) ExecuteCommand() string {
	return fmt.Sprintf("./%s", r.ExecutableFileName())
}

func init() {
	registry[config.LangC] = &_CRunner{}
}
