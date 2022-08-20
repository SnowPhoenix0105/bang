package hack

import (
	gerrors "github.com/pkg/errors"
)

func As(err, target error) {
	gerrors.As(err, target)
}
