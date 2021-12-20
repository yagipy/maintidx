package maintidx_test

import (
	"testing"

	"github.com/yagipy/maintidx"

	"golang.org/x/tools/go/analysis/analysistest"
)

func init() {
	maintidx.Analyzer.Flags.Set("under", "100")
}

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, maintidx.Analyzer, "maintidx")
}
