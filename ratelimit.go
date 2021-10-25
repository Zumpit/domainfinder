package domainfinder

import (
    "errors"
    "golang.org/x/time/rate"
)

var ErrBlocked = errors.New("google block")
var RateLimit = rate.NewLimiter(rate.Inf, 0)
