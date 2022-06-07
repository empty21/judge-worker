package runner

func registerPascalRunner() {
	ListRunner = append(ListRunner, runner{
		Code:           "PASCAL",
		SourceFileName: "main.pas",
		ExecFileName:   "main",
		CompileCommand: "fpc main.pas",
		ExecCommand:    "./main",
	})
}
