UNAME := $(shell uname -s)

CXX = clang++
CXX_FLAGS += --std=c++17 -g -O3 -I../

ifneq ($(UNAME), Darwin)
	LD_FLAGS += -lcrypto
endif

LIBAOC_HDRS = $(wildcard ../lib/*.h)
LIBAOC_SRCS = $(wildcard ../lib/*.cc)
LIBAOC_OBJS = $(patsubst ../lib/%.cc, %.o, $(LIBAOC_SRCS))

SOLUTION_SRCS = $(wildcard day*.cc)
SOLUTION_OBJS = $(patsubst %.cc, %.o, $(SOLUTION_SRCS))

BIN = aoc2022
BIN_SRCS = $(LIBAOC_HDRS) $(LIBAOC_OBJS) $(SOLUTION_OBJS)
BIN_DEPS = $(LIBAOC_OBJS) $(SOLUTION_OBJS)

all: $(BIN)

%.o: %.cc
	$(CXX) $(CXX_FLAGS) -o $@ -c $<

#
# libaoc
#

%.o: ../lib/%.cc
	$(CXX) $(CXX_FLAGS) -o $@ -c $<

testing.o: ../lib/testing.cc
	$(CXX) $(CXX_FLAGS) -o $@ -c -DTEST_MAIN $<

#
# aoc2022
#

$(BIN): $(BIN_SRCS)
	$(CXX) $(LD_FLAGS) -o $@ $(BIN_DEPS)

check: $(BIN)
	./$(BIN)

clean:
	rm -f $(BIN) $(LIBAOC_OBJS) $(SOLUTION_OBJS)

.PHONY: all check clean
