package runner

func registerCPPRunner() {
	ListRunner = append(ListRunner, runner{
		Code:           "CPP",
		SourceFileName: "main.cpp",
		ExecFileName:   "main.o",
		CompileCommand: "g++ -std=c++14 -o main.o main.cpp",
		ExecCommand:    "./main.o",
	})
}
