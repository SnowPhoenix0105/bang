package deepcopy

import "errors"

const DEBUG = true

var (
	ErrConfigNotValid = errors.New("config is set to an invalid value")
)
