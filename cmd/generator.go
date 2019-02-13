package cmd

import (
	"io"

	chglog "github.com/realeyes-media/git-chglog/pkg/chglog"
)

// Generator ...
type Generator interface {
	Generate(io.Writer, string, *chglog.Config) error
}

type generatorImpl struct{}

// NewGenerator ...
func NewGenerator() Generator {
	return &generatorImpl{}
}

// Generate ...
func (*generatorImpl) Generate(w io.Writer, query string, config *chglog.Config) error {
	return chglog.NewGenerator(config).Generate(w, query)
}

type generatorSemVer struct{}

// NewGeneratorSemVer : Returns a SemVer safe generator
func NewGeneratorSemVer() Generator {
	return &generatorSemVer{}
}

// Generate : Same as Generate but uses SemVer save filter
func (*generatorSemVer) Generate(w io.Writer, query string, config *chglog.Config) error {
	return chglog.NewGenerator(config).GenerateSemVer(w, query)
}
