package checker

import (
	"honnef.co/go/tools/lint"
)

// New returns a new lint.Checker.
func New() lint.Checker {
	return &checker{contextNames: []string{"context.Context"}}
}

type checker struct {
	contextNames []string
}

func (*checker) Name() string            { return "gsc" }
func (*checker) Prefix() string          { return "GSC" }
func (*checker) Init(prog *lint.Program) {}

func (c *checker) Funcs() map[string]lint.Func {
	return map[string]lint.Func{
		"CtxScope": c.CheckCtxScope,
	}
}
