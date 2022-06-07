package runner

func registerNodeRunner() {
	ListRunner = append(ListRunner, runner{
		Code:           "NODE",
		SourceFileName: "main.js",
		ExecFileName:   "main.js",
		CompileCommand: "",
		ExecCommand:    "node main.js",
	})
}
