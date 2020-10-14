PYTHON := /usr/bin/env python3
GO := /usr/bin/env go

TESTS := $(addsuffix .test, $(basename $(wildcard day*.py))) day10.test

all: check

check: $(TESTS) lint

%.test: %.py
	python3 -m unittest $<

day10.test: day10/*.go
	cd day10; $(GO) test -v

lint:
	$(PYTHON) -m black . --check --diff --quiet
	$(PYTHON) -m mypy .
