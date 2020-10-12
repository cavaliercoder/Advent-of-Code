PYTHON := /usr/bin/env python3

TESTS := $(addsuffix .test, $(basename $(wildcard day*.py)))

all: $(TESTS)

%.test: %.py
	python3 -m unittest $<
