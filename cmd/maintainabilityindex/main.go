package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"maintainabilityindex"
)

func main() { singlechecker.Main(maintainabilityindex.Analyzer) }
