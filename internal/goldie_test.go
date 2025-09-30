package internal_test

import (
	"testing"

	"github.com/sebdah/goldie/v2"
)

func newGoldie(t *testing.T) *goldie.Goldie {
	return goldie.New(t,
		// goldie.WithDiffEngine(goldie.ColoredDiff),
		goldie.WithTestNameForDir(true),
	)
}
