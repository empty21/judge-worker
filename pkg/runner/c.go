package runner

func registerCRunner() {
	ListRunner = append(ListRunner, runner{
		Code:           "C",
		SourceFileName: "main.c",
		ExecFileName:   "main.o",
		CompileCommand: "gcc -o main.o main.c",
		ExecCommand:    "./main.o",
	})
}
