package judger

import (
	"fmt"
	"io/fs"
	"judger/pkg/domain"
	"judger/pkg/logger"
	"judger/pkg/runner"
	"judger/pkg/sandbox"
	"judger/pkg/sandbox/docker"
	"judger/pkg/util"
	"os"
	"path"
	"time"
)

const FileCacheTTL = 30 * time.Minute

var Judger *judger

type judger struct {
	sandbox sandbox.Sandbox
}

func (j *judger) Judge(judgeTask domain.JudgeTask) domain.JudgeTaskResult {
	_runner := runner.MapperRunner[judgeTask.LanguageCode]
	if _runner == nil {
		logger.Logger.Error("Language not supported: ", judgeTask.LanguageCode)
		return domain.NewJudgeTaskStatus(judgeTask.Uid, domain.TaskStatusIE)
	}
	workDir := fmt.Sprintf("data/task/%v", judgeTask.Uid)
	defer postJudge(judgeTask)
	err := os.MkdirAll(workDir+"/tests", fs.ModeDir)
	if err != nil {
		logger.Logger.Error(err)
		return domain.NewJudgeTaskStatus(judgeTask.Uid, domain.TaskStatusIE)
	}
	// Write source to file
	sourcePath := path.Join(workDir, _runner.GetSourceFileName())
	err = util.WriteToFile(judgeTask.Source, sourcePath)
	if err != nil {
		logger.Logger.Error(err)
		return domain.NewJudgeTaskStatus(judgeTask.Uid, domain.TaskStatusIE)
	}
	if _runner.GetCompileCommand() != "" {
		err, compileExitStatus := j.sandbox.Compile(_runner, workDir)
		if compileExitStatus != 0 {
			logger.Logger.Error(err)
			return domain.NewJudgeTaskStatus(judgeTask.Uid, domain.TaskStatusCE)
		}
	}
	err = pullTestFromUri(judgeTask.Tests)
	if err != nil {
		return domain.NewJudgeTaskStatus(judgeTask.Uid, domain.TaskStatusIE)
	}
	var testResults []domain.TestResult
	for _, test := range judgeTask.Tests {
		err, testResult := j.sandbox.Execute(_runner, workDir, test, judgeTask.TaskLimitation)
		if err != nil {
			logger.Logger.Error(err)
			testResults = append(testResults, domain.TestResult{
				TestUuid: test.Uuid,
				Result:   domain.TestResultRTE,
			})
		} else {
			testResults = append(testResults, *testResult)
		}
	}
	return domain.NewJudgeTaskResult(judgeTask.Uid, testResults)
}

func postJudge(judgeTask domain.JudgeTask) {
	go func() {
		workDir := fmt.Sprintf("data/task/%v", judgeTask.Uid)
		err := util.ZipAndBackup(workDir)
		if err != nil {
			logger.Logger.Error(err)
		}
		err = os.RemoveAll(workDir)
		if err != nil {
			logger.Logger.Error(err)
		}
	}()
}

func pullTestFromUri(tests []domain.Test) error {
	for _, test := range tests {
		if util.FileNotExisted(fmt.Sprintf("data/tests/%v.in", test.Uuid), FileCacheTTL) {
			err := util.DownloadFile(fmt.Sprintf("data/tests/%v.in", test.Uuid), test.InputUri)
			if err != nil {
				return err
			}
		}

		if util.FileNotExisted(fmt.Sprintf("data/tests/%v.out", test.Uuid), FileCacheTTL) {
			err := util.DownloadFile(fmt.Sprintf("data/tests/%v.out", test.Uuid), test.OutputUri)

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func init() {
	_ = os.MkdirAll("data/tests", fs.ModeDir)
	Judger = &judger{sandbox: docker.NewDockerSandbox("empty21/judge-sandbox")}
}
