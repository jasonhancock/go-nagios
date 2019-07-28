all: test

deps:
	go get github.com/stretchr/testify
	go get github.com/pkg/errors
test:
	go test -v
