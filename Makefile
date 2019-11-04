APP             = pub
PACKAGE  		= github.com/devigned/$(APP)
DATE    		?= $(shell date +%FT%T%z)
VERSION 		?= $(shell git rev-list -1 HEAD)
SHORT_VERSION 	?= $(shell git rev-parse --short HEAD)
GOBIN      		?= $(HOME)/go/bin
GOFMT   		= gofmt
GO      		= go
PKGS     		= $(or $(PKG),$(shell $(GO) list ./... | grep -vE "^$(PACKAGE)/templates/"))

V = 0
Q = $(if $(filter 1,$V),,@)

.PHONY: all
all: fmt lint vet tidy build


GOLINT = $(GOBIN)/golint
$(GOBIN)/golint: ; $(info $(M) building golint…)
	$(GO) get -u golang.org/x/lint/golint

build: lint tidy ; $(info $(M) buiding ./bin/pub)
	$Q $(GO)  build -ldflags "-X $(PACKAGE)/cmd.GitCommit=$(VERSION)" -o ./bin/$(APP)

.PHONY: lint
lint: $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q test -z "$$($(GOLINT) ./... | tee /dev/stderr)" || exit 1

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./...); do \
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

.PHONY: vet
vet: $(GOLINT) ; $(info $(M) running vet…) @ ## Run vet
	$Q $(GO) vet ./...

.PHONY: tidy
tidy: ; $(info $(M) running tidy…) @ ## Run tidy
	$Q $(GO) mod tidy

.PHONY: build-debug
build-debug: ; $(info $(M) buiding debug...)
	$Q $(GO)  build -o ./bin/$(APP) -tags debug

.PHONY: gox
gox:
	gox -osarch="darwin/amd64 windows/amd64 linux/amd64" -ldflags "-X $(PACKAGE)/cmd.GitCommit=$(VERSION)" -output "./bin/$(SHORT_VERSION)/{{.Dir}}_{{.OS}}_{{.Arch}}"
