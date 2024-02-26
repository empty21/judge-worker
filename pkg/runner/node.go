package runner

import (
	"fmt"
	"judger/pkg/config"
)

type _NodeRunner struct {
}

func (r *_NodeRunner) SandboxImage() string {
	return config.SandboxImageNode
}

func (r *_NodeRunner) SourceFileName() string {
	return "main.js"
}

func (r *_NodeRunner) ExecutableFileName() string {
	return r.SourceFileName()
}

func (r *_NodeRunner) CompileCommand() string {
	return ""
}

func (r *_NodeRunner) ExecuteCommand() string {
	return fmt.Sprintf("node %s", r.ExecutableFileName())
}

func init() {
	registry[config.LangNode] = &_NodeRunner{}
}
