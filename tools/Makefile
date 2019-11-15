GO 		?= go
GOBIN	?= $(shell $(GO) env GOPATH)/bin
GOLINT	= $(GOBIN)/golint

default: all

all: $(GOBIN)/golint $(GOBIN)/goveralls
	@echo > /dev/null

$(GOBIN)/golint:
	$(GO) get golang.org/x/lint/golint

$(GOBIN)/goveralls:
	$(GO) get github.com/mattn/goveralls

