// https://golangci-lint.run/contributing/new-linters/#how-to-add-a-private-linter-to-golangci-lint
package main

import (
	"strings"

	"github.com/yagipy/maintidx"
	"golang.org/x/tools/go/analysis"
)

var flags string

// AnalyzerPlugin provides analyzers as a plugin.
// It follows golangci-lint style plugin.
var AnalyzerPlugin analyzerPlugin

type analyzerPlugin struct{}

func (analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	if flags != "" {
		flagset := maintidx.Analyzer.Flags
		if err := flagset.Parse(strings.Split(flags, " ")); err != nil {
			panic("cannot parse flags of maintidx: " + err.Error())
		}
	}
	return []*analysis.Analyzer{
		maintidx.Analyzer,
	}
}
