package runner

import (
	"fmt"
	"judger/pkg/config"
	"strings"
)

type _RustRunner struct {
}

func (r *_RustRunner) SandboxImage() string {
	return config.SandboxImageRust
}

func (r *_RustRunner) SourceFileName() string {
	return "main.rb"
}

func (r *_RustRunner) ExecutableFileName() string {
	return strings.Replace(r.SourceFileName(), ".rb", ".o", -1)
}

func (r *_RustRunner) CompileCommand() string {
	return fmt.Sprintf("rustc -o %s %s", r.ExecutableFileName(), r.SourceFileName())
}

func (r *_RustRunner) ExecuteCommand() string {
	return fmt.Sprintf("./%s", r.ExecutableFileName())
}

func init() {
	registry[config.LangRust] = &_RustRunner{}
}
