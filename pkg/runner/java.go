package runner

func registerJavaRunner() {
	ListRunner = append(ListRunner, runner{
		Code:           "JAVA",
		SourceFileName: "main.java",
		ExecFileName:   "Execute",
		CompileCommand: "javac main.java",
		ExecCommand:    "java Execute",
	})
}
