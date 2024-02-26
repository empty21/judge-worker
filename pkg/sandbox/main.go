package sandbox

import (
	"errors"
)

var registry = make(map[string]Sandbox)

func GetSandbox(name string) (Sandbox, error) {
	if registry[name] == nil {
		return nil, errors.New("sandbox not found")
	}
	return registry[name], nil
}
