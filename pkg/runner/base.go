package runner

type Runner interface {
	SandboxImage() string
	SourceFileName() string
	ExecutableFileName() string
	CompileCommand() string
	ExecuteCommand() string
}
