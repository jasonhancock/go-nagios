all: test

deps:
	go get github.com/cheekybits/is
	go get github.com/pkg/errors
test:
	go test -v
