DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)

test: test-deps
	go list ./... | xargs -n1 go test

test-deps:
	go get github.com/stretchr/testify

release-deps:
	go get github.com/mitchellh/gox
