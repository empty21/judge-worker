package sandbox

import (
	"fmt"
	"judger/pkg/config"
	"judger/pkg/log"
	"judger/pkg/runner"
	"judger/pkg/util"
	"os"
	"os/exec"
	"path"
	"strings"
)

type docker struct {
}

func (d *docker) Exists() bool {
	command := "docker"

	args := []string{"info"}
	cmd := exec.Command(command, args...)
	err, _ := util.RunCommand(cmd)
	return err == nil
}

func (d *docker) Compile(runner runner.Runner, workSpace string) int {
	command := "docker"

	args := []string{
		`run`,
		`-i`,
		`--rm`,
		`-v`,
		fmt.Sprintf("%s:%s", config.GetJudgeVolume(), config.AbsoluteJudgeFolder),
		`-w`,
		strings.Replace(workSpace, config.JudgeFolder, config.AbsoluteJudgeFolder, 1),
		runner.SandboxImage(),
	}
	args = append(args, strings.Split(runner.CompileCommand(), " ")...)
	outputFile, _ := os.Create(path.Join(workSpace, config.FileNameCompileLog))
	errorFile, _ := os.Create(path.Join(workSpace, config.FileNameCompileError))
	defer func() {
		_ = outputFile.Close()
		_ = errorFile.Close()
	}()
	cmd := exec.Command(command, args...)

	cmd.Stdout = outputFile
	cmd.Stderr = errorFile
	_, exitCode := util.RunCommand(cmd)

	return exitCode
}

func (d *docker) Execute(runner runner.Runner, workSpace string, test string, option ExecuteOption) int {
	command := "docker"

	args := []string{
		`run`,
		`-a`,
		`stdin`,
		`-a`,
		`stdout`,
		`cpus=1`,
		fmt.Sprintf(`--memory=%vm`, option.MemoryLimit),
		`-i`,
		`--rm`,
		`-v`,
		fmt.Sprintf("%s:%s", config.GetJudgeVolume(), config.AbsoluteJudgeFolder),
		`-w`,
		strings.Replace(workSpace, config.JudgeFolder, config.AbsoluteJudgeFolder, 1),
		runner.SandboxImage(),
		"sh",
		"-c",
	}
	shCommand := fmt.Sprintf(
		"time -f %%e:%%M -o %v timeout %v %v < %v > %v 2> %v",
		path.Join(config.JudgeFolderTests, fmt.Sprintf(config.FileNameTemplateStat, test)),
		option.TimeLimit,
		runner.ExecuteCommand(),
		path.Join(config.JudgeFolderTests, fmt.Sprintf(config.FileNameTemplateInput, test)),
		path.Join(config.JudgeFolderTests, fmt.Sprintf(config.FileNameTemplateOutput, test)),
		path.Join(config.JudgeFolderTests, fmt.Sprintf(config.FileNameTemplateError, test)),
	)
	args = append(args, shCommand)
	cmd := exec.Command(command, args...)
	err, exitCode := util.RunCommand(cmd)
	if err != nil {
		log.Error("Execute error: %v", err.Error())
	}

	return exitCode
}

func init() {
	registry[config.SandboxDocker] = &docker{}
}
