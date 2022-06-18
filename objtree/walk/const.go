package walk

import "errors"

var (
	ErrCallMethodOnInvalidWalkNodeContext = errors.New("call method on an invalid WalkNodeContext")
)
