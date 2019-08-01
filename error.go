package round

import (
    "errors"
)

var(
    ErrNoServerAlive    = errors.New("There is no service Alive")
    ErrInvalidWeight    = errors.New("Invalid Server weight parameter")
)

