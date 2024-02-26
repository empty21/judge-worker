package config

import (
	"fmt"
	"os"
)

const (
	TaskQueueName   = "judge.tasks.queue"
	ResultQueueName = "judge.results.queue"
)

const (
	FileNameCompileLog             = "compile.log"
	FileNameCompileError           = "compile.err"
	FileNameTemplateInput          = "%s.in"
	FileNameTemplateOutput         = "%s.out"
	FileNameTemplateExpectedOutput = "%s.expected.out"
	FileNameTemplateError          = "%s.err"
	FileNameTemplateStat           = "%s.stat"
	FilenameTemplateZip            = "%s.zip"
)

const CacheFolder = ".cache"
const JudgeFolder = ".judge"
const AbsoluteJudgeFolder = "/" + JudgeFolder
const JudgeFolderWorkspaces = "workspaces"
const JudgeFolderTests = "tests"
const FileCachedCount = 1000

func GetJudgeVolume() string {
	if Config.Debug {
		wd, _ := os.Getwd()
		return fmt.Sprintf("%s/%s", wd, JudgeFolder)
	}
	return "judge-volume"
}
