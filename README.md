# maintidx
`maintidx` measures the maintainability index of each function.  
[Here](https://blog.yagipy.me/analyze-maintainability-index) for more information about this package.(Sorry, japanese only)

## What is maintainability index
The maintainability index is an index that measures the maintainability of source code.  
The coefficients include cyclomatic complexity, halstead volume, and line of code.  
This library used [the rebased value provided by Microsoft.](https://docs.microsoft.com/en-us/visualstudio/code-quality/code-metrics-maintainability-index-range-and-meaning)  
The maintainability index is an experimental value, so don't expect too much from it, but [the cyclomatic complexity used as a coefficient was found to be closely related to the number of times the code was edited.](https://ieeexplore.ieee.org/document/312034)

## Installation
### Go version < 1.16
```shell
go get -u github.com/yagipy/maintidx/cmd/maintidx@v1.0.0
```

### Go version 1.16+
```shell
go install github.com/yagipy/maintidx/cmd/maintidx@v1.0.0
```

## Usage
### standalone
```shell
maintidx ./...
```

### with [golangci-lint](https://github.com/golangci/golangci-lint)
1. [Install golangci-lint](https://golangci-lint.run/usage/install/)
2. Execute the following
```shell
golangci-lint run --disable-all -E maintidx
```

It's also available to use .golangci.yml
```yaml
# Add maintidx to enable linters.
linters:
  enable:
    - maintidx

linters-settings:
  maintidx:
    # Show functions with maintainability index lower than N.
    # A high index indicates better maintainability (it's kind of the opposite of complexity).
    # Default: 20
    under: 100
```

Execute using .golangci.yml
```shell
golangci-lint run
```

### with go run
```shell
go run github.com/yagipy/maintidx/cmd/maintidx ./...
```

### with go vet
```shell
go vet -vettool=`which maintidx` ./...
```

## Flag
```shell
Flags:
  -under int
    	show functions with maintainability index < N only. (default 20)
```

## TODO
- [ ] Setup execute env on container
- [ ] Impl cyc.Cyc.Calc()
- [ ] Move maintidx.Visitor.PrintHalstVol to halstval package
- [ ] Consider the necessity of halstvol.incrIfAllTrue
- [ ] Test under pkg file
