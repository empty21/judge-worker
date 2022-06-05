package util

import "os/exec"

const (
	DefaultSuccessCode    = 0
	DefaultFailedCode     = 1
	TimeLimitExceededCode = 124
)

func RunCommand(c *exec.Cmd) (error, int) {
	err := c.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return err, exitErr.ExitCode()
		}
		return err, DefaultFailedCode
	}

	return nil, DefaultSuccessCode
}
