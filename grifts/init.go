package grifts

import (
	"github.com/akosgarai/buffalo_example/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
