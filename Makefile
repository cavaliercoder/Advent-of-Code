all: check

check:
	cd go/; make
	cd python/; make
