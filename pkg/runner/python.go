package runner

import (
	"fmt"
	"judger/pkg/config"
)

type _PythonRunner struct {
}

func (r *_PythonRunner) SandboxImage() string {
	return config.SandboxImagePython
}

func (r *_PythonRunner) SourceFileName() string {
	return "main.py"
}

func (r *_PythonRunner) ExecutableFileName() string {
	return r.SourceFileName()
}

func (r *_PythonRunner) CompileCommand() string {
	return ""
}

func (r *_PythonRunner) ExecuteCommand() string {
	return fmt.Sprintf("python %s", r.ExecutableFileName())
}

func init() {
	registry[config.LangPython] = &_PythonRunner{}
}
