package test

import (
	"github.com/onsi/gomega"
	"testing"

	"kitchen/pkg/core"
)

func getTestContext() *core.Context {
	ctx, _ := core.NewContext("")
	ctx.LogConfig.File = ""
	ctx.LogConfig.Level = "debug"
	_ = ctx.UpdateLogFileSettings()
	return ctx
}

type Test struct {
	*gomega.GomegaWithT
	*core.Context
	tt *testing.T
}

func (t *Test) Run(name string, f func()) {
	t.tt.Run(name, func(tt *testing.T) {
		g := gomega.NewGomegaWithT(tt)
		savedT := t.GomegaWithT
		t.GomegaWithT = g
		f()
		t.GomegaWithT = savedT
	})
}

func NewTest(t *testing.T) *Test {
	g := gomega.NewGomegaWithT(t)
	gomega.RegisterTestingT(t)
	bt := &Test{
		GomegaWithT: g,
		tt:          t,
		Context:     getTestContext(),
	}
	return bt
}

func NewTestWithoutContext(t *testing.T) *Test {
	g := gomega.NewGomegaWithT(t)
	gomega.RegisterTestingT(t)
	bt := &Test{
		GomegaWithT: g,
		tt:          t,
	}
	return bt
}
