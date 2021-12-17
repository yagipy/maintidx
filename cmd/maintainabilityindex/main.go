package main

import (
	"maintainabilityindex"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(maintainabilityindex.Analyzer) }
