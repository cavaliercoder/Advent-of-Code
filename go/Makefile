GO := /usr/bin/env go

TESTS := $(addsuffix .test, $(wildcard day*))

all: check

check: $(TESTS)

%.test: %/*.go
	cd $(basename $@)/; $(GO) test -v
