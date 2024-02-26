package util

import (
	"errors"
	"judger/pkg/config"
	"os/exec"
)

func RunCommand(c *exec.Cmd) (error, int) {
	err := c.Run()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return err, exitErr.ExitCode()
		}
		return err, config.ExitCodeError
	}

	return nil, config.ExitCodeSuccess
}
