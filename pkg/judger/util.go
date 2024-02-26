package judger

import (
	"fmt"
	"judger/pkg/config"
	"judger/pkg/file"
	"judger/pkg/log"
	"judger/pkg/model"
	"path"
	"strconv"
	"strings"
)

// Extract time and memory from stat file
func extractStat(filePath string) (float64, int64) {
	fileContent, _ := file.Read(filePath)
	s := strings.Split(strings.Trim(fileContent, " \n\r\t"), ":")
	if len(s) == 2 {
		time, _ := strconv.ParseFloat(s[0], 64)
		memory, _ := strconv.ParseInt(s[1], 10, 64)
		return time, memory
	}
	return 0, 0
}

func prepareTestFiles(workSpace string, tests []model.Test) error {
	for _, test := range tests {
		inputPath := path.Join(workSpace, config.JudgeFolderTests, fmt.Sprintf(config.FileNameTemplateInput, test.Identifier))
		err := file.Cache(file.GetFile)(inputPath, test.InputUri)
		if err != nil {
			return err
		}
		outputPath := path.Join(workSpace, config.JudgeFolderTests, fmt.Sprintf(config.FileNameTemplateExpectedOutput, test.Identifier))
		err = file.Cache(file.GetFile)(outputPath, test.OutputUri)
		if err != nil {
			return err
		}
	}
	return nil
}

func judgeResult(workSpace string, test model.Test, exitCode int) model.TestResult {
	time, memory := extractStat(path.Join(workSpace, config.JudgeFolderTests, fmt.Sprintf(config.FileNameTemplateStat, test.Identifier)))

	testResult := model.TestResult{
		Identifier: test.Identifier,
		Result:     config.TestResultRTE,
		Memory:     memory,
		Time:       time,
	}

	if exitCode == config.ExitCodeSuccess {
		outputPath := path.Join(workSpace, config.JudgeFolderTests, fmt.Sprintf(config.FileNameTemplateOutput, test.Identifier))
		expectedOutputPath := path.Join(workSpace, config.JudgeFolderTests, fmt.Sprintf(config.FileNameTemplateExpectedOutput, test.Identifier))
		output, _ := file.Read(outputPath)
		expectedOutput, _ := file.Read(expectedOutputPath)
		if strings.Compare(normalizeLineBreaker(output), normalizeLineBreaker(expectedOutput)) == 0 {
			testResult.Result = config.TestResultAC
		} else {
			testResult.Result = config.TestResultWA
		}
	}

	if exitCode == config.ExitCodeTLE {
		testResult.Result = config.TestResultTLE
	}

	return testResult
}

func normalizeLineBreaker(s string) string {
	replacer := strings.NewReplacer("\r\n", "\n", "\r", "\n")
	return replacer.Replace(s)
}

func cleanUp(workSpace string) {
	backUp(workSpace)
	_ = file.Remove(workSpace)
}

func backUp(workSpace string) {
	workSpaceIdentifier := path.Base(workSpace)
	zipPath := path.Join(config.JudgeFolder, fmt.Sprintf(config.FilenameTemplateZip, workSpaceIdentifier))
	err := file.Compress(zipPath, workSpace)
	if err != nil {
		log.Error("Backup workspace error: %v", err.Error())
	}
	if err != nil {
		log.Error("Upload backup error: %v", err.Error())
	} else {
		log.Info("Backup workspace: %v", workSpaceIdentifier)
	}
}
