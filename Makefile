.PHONY: build

build:
	rm -rf dist || true
	mkdir dist
	cd src/vm && gcc lc3.c -o lc3-vm
	mv src/vm/lc3-vm dist
