package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"maintidx"
)

func main() { singlechecker.Main(maintidx.Analyzer) }
