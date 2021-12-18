package maintainabilityindex_test

import (
	"github.com/gostaticanalysis/testutil"
	"testing"

	"maintainabilityindex"

	"golang.org/x/tools/go/analysis/analysistest"
)

func init() {
	maintainabilityindex.Analyzer.Flags.Set("under", "100")
}

func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, maintainabilityindex.Analyzer, "maintainabilityindex")
}
