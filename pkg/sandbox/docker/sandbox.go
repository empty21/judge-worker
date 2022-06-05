package docker

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"judger/pkg/domain"
	"judger/pkg/runner"
	"judger/pkg/sandbox"
	"judger/pkg/util"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

type dockerSandbox struct {
}

func (d dockerSandbox) Compile(runner runner.Runner, workDir string) (error, int) {
	command := "docker"
	args := []string{
		`run`,
		`-i`,
		`--rm`,
		`-v`,
		fmt.Sprintf("%v:%v", getAbsolutePath(workDir), getPathInDocker(workDir)),
		`-w`,
		getPathInDocker(workDir),
		`empty21/judge-sandbox`,
	}
	args = append(args, strings.Split(runner.GetCompileCommand(), " ")...)
	outputFile, _ := os.Create(path.Join(workDir, "compile.out"))
	errorFile, _ := os.Create(path.Join(workDir, "compile.err"))
	defer func() {
		_ = outputFile.Close()
		_ = errorFile.Close()
	}()
	cmd := exec.Command(command, args...)
	cmd.Stdout = outputFile
	cmd.Stderr = errorFile
	err, statusCode := util.RunCommand(cmd)
	if err != nil {
		return err, util.DefaultFailedCode
	}
	return err, statusCode
}

func (d dockerSandbox) Execute(runner runner.Runner, workDir string, test domain.Test, limitation domain.TaskLimitation) (error, *domain.TestResult) {
	var result = &domain.TestResult{
		TestUuid: test.Uuid,
		Memory:   0,
		Time:     0,
	}

	command := "docker"
	args := []string{
		`run`,
		`-a`,
		`stdin`,
		`-a`,
		`stdout`,
		fmt.Sprintf(`--memory=%vm`, limitation.MemoryLimit),
		`-i`,
		`--rm`,
		`-v`,
		fmt.Sprintf(`%v:%v`, getAbsolutePath(workDir), getPathInDocker(workDir)),
		`-w`,
		getPathInDocker(workDir),
		`empty21/judge-sandbox`,
		`/usr/bin/time`,
		`-q`,
		`-f`,
		`stat:%e:%M`,
		`-o`,
		fmt.Sprintf(`./tests/%v.log`, test.Uuid),
		`timeout`,
		fmt.Sprintf(`%v`, limitation.TimeLimit),
	}
	args = append(args, strings.Split(runner.GetExecCommand(), " ")...)
	outputFile, _ := os.Create(path.Join(workDir, "/tests/"+test.Uuid+".out"))
	errorFile, _ := os.Create(path.Join(workDir, "/tests/"+test.Uuid+".err"))
	cmd := exec.Command(command, args...)
	cmd.Stdout = outputFile
	cmd.Stderr = errorFile
	inputFile, err := os.Open(fmt.Sprintf("data/tests/%v.in", test.Uuid))
	if err != nil {
		return err, nil
	}
	cmd.Stdin = inputFile

	defer func() {
		_ = inputFile.Close()
		_ = outputFile.Close()
		_ = errorFile.Close()
	}()
	err, exitCode := util.RunCommand(cmd)
	if err != nil {
		return err, nil
	}

	// Judge
	judgeByExitedCode(exitCode, result)
	if result.Result == "" {
		judgeByOutput(workDir, result)
	}
	getExecStat(workDir, result)

	return nil, result
}

func judgeByExitedCode(exitCode int, t *domain.TestResult) {
	if exitCode == util.DefaultSuccessCode {
		return
	}
	if exitCode == util.TimeLimitExceededCode {
		t.Result = domain.TestResultTLE
		return
	}
	t.Result = domain.TestResultRTE
	return
}

func judgeByOutput(workDir string, t *domain.TestResult) {
	actualOutput, err := ioutil.ReadFile(fmt.Sprintf("%v/tests/%v.out", workDir, t.TestUuid))
	if err != nil {
		return
	}
	expectedOutput, err := ioutil.ReadFile(fmt.Sprintf("data/tests/%v.out", t.TestUuid))
	if err != nil {
		return
	}
	if bytes.Compare(bytes.TrimSpace(actualOutput), bytes.TrimSpace(expectedOutput)) == 0 {
		t.Result = domain.TestResultAC
	} else {
		t.Result = domain.TestResultWA
	}
}

func getExecStat(workDir string, t *domain.TestResult) {
	file, err := os.Open(fmt.Sprintf("%v/tests/%v.log", workDir, t.TestUuid))
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "stat") == 0 {
			result := strings.Split(line, ":")
			if len(result) == 3 {
				time, _ := strconv.ParseFloat(result[1], 64)
				t.Time = time
				memory, _ := strconv.ParseInt(result[2], 10, 64)
				t.Memory = memory
			}
			return
		}
	}
}

func getPathInDocker(p string) string {
	return strings.Replace(p, "data", "/data", 1)
}

func getAbsolutePath(p string) string {
	wd, err := os.Getwd()
	if err != nil {
		return p
	}
	return path.Join(wd, p)
}

func NewSandbox() sandbox.Sandbox {
	return &dockerSandbox{}
}
