api:	$(shell find app bizrules cmd -type f -name "*.go")
	go build -o api ./cmd/api/main.go

.PHONY: import-paths
import-paths:
	for path in $(shell go list ./...); do    \
	  cd "$(GOPATH)/src/$$path" &&            \
	  FORMAT=svg FILTER=dot gen-import-path;  \
	done

.PHONY: test
test:
	go test -v $(shell go list ./... | grep -v /vendor/)
