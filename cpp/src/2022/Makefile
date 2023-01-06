SOLUTIONS = $(patsubst %.cc, %.o, $(wildcard day*.cc))

LIBAOC = aoc.o
LIBAOC_SOURCES = ../lib/aoc.cc
LIBAOC_HEADERS = ../lib/aoc.h

BIN = aoc2022
BIN_SOURCES = aoc2022.cc

CXX = clang++
CXX_FLAGS = --std=c++17 -I../

all: $(BIN)

%.o: %.cc $(LIBAOC_HEADERS)
	$(CXX) $(CXX_FLAGS) -o $@ -c $<

$(LIBAOC): $(LIBAOC_SOURCES) $(LIBAOC_HEADERS)
	$(CXX) $(CXX_FLAGS) -c -o $@ $(LIBAOC_SOURCES)

$(BIN): $(BIN_SOURCES) $(LIBAOC) $(SOLUTIONS)
	$(CXX) $(CXX_FLAGS) -o $@ $^

check: $(BIN)
	./$(BIN)

clean:
	rm -f $(BIN) $(LIBAOC) $(SOLUTIONS)

.PHONY: all check clean
