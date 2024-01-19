package errors

import "errors"

var (
	ErrMultipleBots = errors.New("multiple bots cannot play a game")
)
