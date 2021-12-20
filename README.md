# maintidx
maintidx measures the maintainability index of each function.

## Generate exec file
```shell
make build
```

## Run test
```shell
go test ./...
```

## TODO
- [ ] Setup execute env on container
- [ ] Impl cyc.Cyc.Calc()
- [ ] Move maintidx.Visitor.PrintHalstVol to halstval package
- [ ] Consider the necessity of halstvol.incrIfAllTrue
- [ ] Test under pkg file
