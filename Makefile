.PHONY: run test tests debug_tests

run:
	@go run .

test: tests

tests:
	@go test ./...

debug_tests:
	@go test -v ./...
