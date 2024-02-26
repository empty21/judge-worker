package runner

import (
	"errors"
	"fmt"
)

var registry = make(map[string]Runner)

func GetRunner(code string) (Runner, error) {
	if registry[code] == nil {
		return nil, errors.New(fmt.Sprintf("Runner for language %s not found", code))
	}
	return registry[code], nil
}
