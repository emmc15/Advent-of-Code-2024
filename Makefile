.PHONY: build
build:
	@echo "Building..."
	@go build -o bin/aoc ./main.go


.PHONY: run
run:
	@echo "Running AOC 2024..."
	@go run .