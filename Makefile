.PHONY:
.SILENT:
.DEFAULT_GOAL := run

run:
	@go run ./cmd/app/main.go
