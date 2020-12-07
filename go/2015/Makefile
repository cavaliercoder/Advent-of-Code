GO := /usr/bin/env go

TESTS := $(addsuffix .test, $(patsubst %/,%,$(wildcard */)))

all: check

check: $(TESTS)

%.test: %/*.go
	cd $(basename $@)/; $(GO) test -v -count=1 -timeout 10s
