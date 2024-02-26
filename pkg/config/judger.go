package config

const (
	LangC      = "C"
	LangCpp    = "CPP"
	LangGo     = "GO"
	LangJava   = "JAVA"
	LangNode   = "NODE"
	LangPascal = "PASCAL"
	LangPython = "PYTHON"
	LangRuby   = "RUBY"
	LangRust   = "RUST"
)

const (
	SandboxImageC      = "quay.io/stupd/sandbox/c"
	SandboxImageCPP    = "quay.io/stupd/sandbox/cpp"
	SandboxImageGo     = "quay.io/stupd/sandbox/go"
	SandboxImageJava   = "quay.io/stupd/sandbox/java"
	SandboxImageNode   = "quay.io/stupd/sandbox/node"
	SandboxImagePascal = "quay.io/stupd/sandbox/pascal"
	SandboxImagePython = "quay.io/stupd/sandbox/python"
	SandboxImageRuby   = "quay.io/stupd/sandbox/ruby"
	SandboxImageRust   = "quay.io/stupd/sandbox/rust"
)

const (
	SandboxDocker = "docker"
	SandboxPodman = "podman"
)

type TestResultEnum string
type TaskStatusEnum string

const (
	TaskStatusIP TaskStatusEnum = "IP"
	TaskStatusCE                = "CE"
	TaskStatusIE                = "IE"
)

const (
	TestResultAC  TestResultEnum = "AC"
	TestResultWA                 = "WA"
	TestResultRTE                = "RTE"
	TestResultTLE                = "TLE"
)

const (
	ExitCodeSuccess = 0
	ExitCodeError   = 1
	ExitCodeTLE     = 124
)
