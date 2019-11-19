GO 			?= go
LOCALBIN 	= $(shell pwd)/bin

default: install

install: $(LOCALBIN)/golint $(LOCALBIN)/goveralls $(LOCALBIN)/gox
	@echo > /dev/null

$(LOCALBIN)/golint:
	GOBIN=$(LOCALBIN) $(GO) get golang.org/x/lint/golint

$(LOCALBIN)/goveralls:
	GOBIN=$(LOCALBIN) $(GO) get github.com/mattn/goveralls

$(LOCALBIN)/gox:
	GOBIN=$(LOCALBIN) $(GO) get github.com/mitchellh/gox

.PHONY: reload
reload: clean install

.PHONY: clean
clean:
	rm -rf $(LOCALBIN)
