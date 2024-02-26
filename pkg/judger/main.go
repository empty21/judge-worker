package judger

import (
	"io/fs"
	"judger/pkg/config"
	"judger/pkg/file"
	"judger/pkg/log"
	"judger/pkg/model"
	"judger/pkg/runner"
	"judger/pkg/sandbox"
	"os"
	"path"
)

type judger struct {
	sandbox sandbox.Sandbox
}

var Judger *judger

func (j *judger) Judge(judgeTask model.JudgeTask) model.JudgeTaskResult {
	r, err := runner.GetRunner(judgeTask.LanguageCode)
	if err != nil {
		log.Error("Get runner error #%v: %v", judgeTask.Identifier, err.Error())
		return model.NewJudgeTaskStatus(judgeTask.Identifier, config.TaskStatusIE, err.Error())
	}
	workSpace := path.Join(config.JudgeFolder, config.JudgeFolderWorkspaces, judgeTask.Identifier)
	defer cleanUp(workSpace)

	err = os.MkdirAll(path.Join(workSpace, config.JudgeFolderTests), fs.ModeDir)
	if err != nil {
		log.Error("Prepare workspace error #%v: %v", judgeTask.Identifier, err.Error())
		return model.NewJudgeTaskStatus(judgeTask.Identifier, config.TaskStatusIE, err.Error())
	}

	sourcePath := path.Join(workSpace, r.SourceFileName())
	err = file.Write(sourcePath, judgeTask.Source)
	if err != nil {
		log.Error("Write source error #%v: %v", judgeTask.Identifier, err.Error())
		return model.NewJudgeTaskStatus(judgeTask.Identifier, config.TaskStatusIE, err.Error())
	}

	if r.CompileCommand() != "" {
		compileExitCode := j.sandbox.Compile(r, workSpace)
		if compileExitCode != config.ExitCodeSuccess {
			compileError, _ := os.ReadFile(path.Join(workSpace, config.FileNameCompileError))
			log.Error("Compile error #%v: %v", judgeTask.Identifier, string(compileError))
			return model.NewJudgeTaskStatus(judgeTask.Identifier, config.TaskStatusCE, string(compileError))
		}
	}

	err = prepareTestFiles(workSpace, judgeTask.Tests)
	if err != nil {
		log.Error("Prepare test files error #%v: %v", judgeTask.Identifier, err.Error())
		return model.NewJudgeTaskStatus(judgeTask.Identifier, config.TaskStatusIE, err.Error())
	}

	var testResults []model.TestResult
	for _, test := range judgeTask.Tests {
		// Copy input file to workspace
		timeLimit := test.TimeLimit
		if timeLimit == 0 {
			timeLimit = judgeTask.TimeLimit
		}
		memoryLimit := test.MemoryLimit
		if memoryLimit == 0 {
			memoryLimit = judgeTask.MemoryLimit
		}

		exitCode := j.sandbox.Execute(r, workSpace, test.Identifier, sandbox.ExecuteOption{
			MemoryLimit: memoryLimit,
			TimeLimit:   timeLimit,
		})

		testResults = append(testResults, judgeResult(workSpace, test, exitCode))
	}
	return model.NewJudgeTaskResult(judgeTask.Identifier, testResults)
}

func init() {
	_ = os.MkdirAll(config.JudgeFolder, fs.ModeDir)
	s, err := sandbox.GetSandbox(config.Config.Sandbox)
	if err != nil {
		log.Error("Get sandbox error: %v", err.Error())
		os.Exit(1)
	}
	if !s.Exists() {
		log.Error("Sandbox %s is not available on machine", config.Config.Sandbox)
		os.Exit(1)
	}
	log.Info("Sandbox %s is ready to use", config.Config.Sandbox)
	Judger = &judger{s}
}
