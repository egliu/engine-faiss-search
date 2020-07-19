GO=go

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

VERSION := v2.2.0
BUILD := `git rev-parse --short=7 HEAD`
TARGETS := engine-faiss-search-master
ALL_TARGETS := $(TARGETS)
project=github.com/engine-faiss-search

LDFLAGS += -X "$(project)/version.BuildTS=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "$(project)/version.GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "$(project)/version.Version=$(VERSION)"
LDFLAGS += -X "$(project)/version.GitBranch=$(shell git rev-parse --abbrev-ref HEAD)"

baseImage=golang:1.12.17

build: $(TARGETS)

$(TARGETS): $(SRC)
	$(GO) build -ldflags '$(LDFLAGS)' $(project)/cmd/$@

clean:
	rm -f $(TARGETS)

dev:
	@baseImage=${baseImage} bash scripts/start-dev.sh