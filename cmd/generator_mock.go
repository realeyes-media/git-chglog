package cmd

import (
	"io"

	chglog "github.com/realeyes-media/git-chglog/pkg/chglog"
)

type mockGeneratorImpl struct {
	ReturnGenerate func(io.Writer, string, *chglog.Config) error
}

func (m *mockGeneratorImpl) Generate(w io.Writer, query string, config *chglog.Config) error {
	return m.ReturnGenerate(w, query, config)
}
