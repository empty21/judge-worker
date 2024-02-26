package runner

import (
	"fmt"
	"judger/pkg/config"
	"strings"
)

type _CPPRunner struct {
}

func (r *_CPPRunner) SandboxImage() string {
	return config.SandboxImageCPP
}

func (r *_CPPRunner) SourceFileName() string {
	return "main.cpp"
}

func (r *_CPPRunner) ExecutableFileName() string {
	return strings.Replace(r.SourceFileName(), ".cpp", ".o", -1)
}

func (r *_CPPRunner) CompileCommand() string {
	return fmt.Sprintf("g++ -std=c++14 -o %s %s", r.ExecutableFileName(), r.SourceFileName())
}

func (r *_CPPRunner) ExecuteCommand() string {
	return fmt.Sprintf("./%s", r.ExecutableFileName())
}

func init() {
	registry[config.LangCpp] = &_CPPRunner{}
}
