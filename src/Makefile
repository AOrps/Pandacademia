.SHELL = /bin/bash

.PHONY = usage compile run clean

all = usage

usage:
	@echo "Usage:"
	@echo "make compile -> Compile file 'main.go' into 'main' executable"
	@echo "make run -> Re-compiles executables and runs it"
	@echo "make clean -> Removes an unnecessary executables"

compile:
	@echo "Compiling main.go into executable"
	@go build main.go

run: compile
	@echo "Running ./main"
	@./main

clean:
ifneq (,$(wildcard ./main))
	@rm ./main 
endif
ifneq (, $(wildcard ./a.out))
	@rm ./a.out
endif