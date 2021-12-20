package maintidx_test

import (
	"github.com/gostaticanalysis/testutil"
	"testing"

	"maintidx"

	"golang.org/x/tools/go/analysis/analysistest"
)

func init() {
	maintidx.Analyzer.Flags.Set("under", "100")
}

func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, maintidx.Analyzer, "maintidx")
}
