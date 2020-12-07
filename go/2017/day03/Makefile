SOURCES = $(wildcard *.go)

all: main

main: $(SOURCES)
	go build -o $@ $^

check: main
	echo 265149 | ./main 
