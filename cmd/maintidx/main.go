package main

import (
	"github.com/yagipy/maintidx"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(maintidx.Analyzer) }
