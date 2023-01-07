CXX = clang++
CXX_FLAGS = --std=c++17 -I../


LIBAOC_HEADERS = $(wildcard ../lib/*.h)
LIBAOC_SOURCES = $(wildcard ../lib/*.cc)
LIBAOC_OBJECTS = $(patsubst ../lib/%.cc, %.o, $(LIBAOC_SOURCES))

SOLUTION_SOURCES = $(wildcard day*.cc)
SOLUTION_OBJECTS = $(patsubst %.cc, %.o, $(SOLUTION_SOURCES))

BIN = aoc2022
BIN_SOURCES = $(LIBAOC_OBJECTS) $(SOLUTION_OBJECTS)

all: $(BIN)

%.o: %.cc
	$(CXX) $(CXX_FLAGS) -o $@ -c $<

%.o: ../lib/%.cc
	$(CXX) $(CXX_FLAGS) -o $@ -c $<

$(BIN): $(BIN_SOURCES)
	$(CXX) $(CXX_FLAGS) -o $@ $^

check: $(BIN)
	./$(BIN)

clean:
	rm -f $(BIN) $(LIBAOC_OBJECTS) $(SOLUTION_OBJECTS)

.PHONY: all check clean
