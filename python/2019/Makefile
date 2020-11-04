PYTHON := /usr/bin/env python3

TESTS := $(addsuffix .test, $(basename $(wildcard day*.py)))

all: check

check: $(TESTS) lint

%.test: %.py
	python3 -m unittest $<

lint:
	$(PYTHON) -m black . --check --diff --quiet
	$(PYTHON) -m mypy .
