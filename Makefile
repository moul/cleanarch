.PHONY: import-paths
import-paths:
	for path in $(shell go list ./...); do    \
	  cd "$(GOPATH)/src/$$path" &&            \
	  FORMAT=svg FILTER=dot gen-import-path;  \
	done
