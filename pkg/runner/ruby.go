package runner

import (
	"fmt"
	"judger/pkg/config"
)

type _RubyRunner struct {
}

func (r *_RubyRunner) SandboxImage() string {
	return config.SandboxImageRuby
}

func (r *_RubyRunner) SourceFileName() string {
	return "main.rb"
}

func (r *_RubyRunner) ExecutableFileName() string {
	return r.SourceFileName()
}

func (r *_RubyRunner) CompileCommand() string {
	return ""
}

func (r *_RubyRunner) ExecuteCommand() string {
	return fmt.Sprintf("ruby %s", r.ExecutableFileName())
}

func init() {
	registry[config.LangRuby] = &_RubyRunner{}
}
