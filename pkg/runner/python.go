package runner

func registerPythonRunner() {
	ListRunner = append(ListRunner, runner{
		Code:           "PYTHON",
		SourceFileName: "main.py",
		ExecFileName:   "main.py",
		CompileCommand: "python3 main.py",
		ExecCommand:    "python3 main.py",
	})
}
