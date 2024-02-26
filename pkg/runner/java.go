package runner

import (
	"fmt"
	"judger/pkg/config"
	"strings"
)

type _JavaRunner struct {
}

func (r *_JavaRunner) SandboxImage() string {
	return config.SandboxImageJava
}

func (r *_JavaRunner) SourceFileName() string {
	return "Execute.java"
}

func (r *_JavaRunner) ExecutableFileName() string {
	return strings.Replace(r.SourceFileName(), ".java", "", -1)
}

func (r *_JavaRunner) CompileCommand() string {
	return fmt.Sprintf("javac %s", r.SourceFileName())
}

func (r *_JavaRunner) ExecuteCommand() string {
	return fmt.Sprintf("java %s", r.ExecutableFileName())
}

func init() {
	registry[config.LangJava] = &_JavaRunner{}
}
