package runner

import (
	"fmt"
	"judger/pkg/config"
	"strings"
)

type _PascalRunner struct {
}

func (r *_PascalRunner) SandboxImage() string {
	return config.SandboxImagePascal
}

func (r *_PascalRunner) SourceFileName() string {
	return "main.pas"
}

func (r *_PascalRunner) ExecutableFileName() string {
	return strings.Replace(r.SourceFileName(), ".pas", ".o", -1)
}

func (r *_PascalRunner) CompileCommand() string {
	return fmt.Sprintf("fpc -o%s %s", r.ExecutableFileName(), r.SourceFileName())
}

func (r *_PascalRunner) ExecuteCommand() string {
	return fmt.Sprintf("./%s", r.ExecutableFileName())
}

func init() {
	registry[config.LangPascal] = &_PascalRunner{}
}
