package runner

type Runner interface {
	GetExecFileName() string
	GetSourceFileName() string
	GetCompileCommand() string
	GetExecCommand() string
}

type runner struct {
	Code           string
	SourceFileName string
	ExecFileName   string
	CompileCommand string
	ExecCommand    string
}

func (r runner) GetExecFileName() string {
	return r.ExecFileName
}

func (r runner) GetSourceFileName() string {
	return r.SourceFileName
}

func (r runner) GetCompileCommand() string {
	return r.CompileCommand
}

func (r runner) GetExecCommand() string {
	return r.ExecCommand
}

var ListRunner []runner

var MapperRunner = make(map[string]Runner)

func init() {
	registerCRunner()
	registerCPPRunner()
	registerJavaRunner()
	registerPythonRunner()
	registerNodeRunner()
	registerPascalRunner()

	for _, r := range ListRunner {
		MapperRunner[r.Code] = r
	}
}
