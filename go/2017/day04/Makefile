SOURCES = $(wildcard *.go)

all: check

main: $(SOURCES)
	go build -o $@ $^

check: main
	./main < ../../../inputs/2017/day04.txt
