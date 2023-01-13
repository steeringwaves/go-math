.DEFAULT_GOAL: all
.PHONY: all tidy test

all: tidy test

test:
	@go test
	@cd clamp && go test
	@cd wrap && go test
	@cd orientation && go test

tidy:
	@go mod tidy
