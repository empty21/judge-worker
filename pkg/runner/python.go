package runner

func registerPythonRunner() {
	ListRunner = append(ListRunner, runner{
		Code:           "PYTHON",
		SourceFileName: "main.py",
		ExecCommand:    "python3 main.py",
	})
}
