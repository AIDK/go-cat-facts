build:
	@echo building to output...
	@go build -o bin/fact cmd/app/main.go
	@echo done

run: build
	@echo running app...
	@./bin/fact

test:
	@echo testing...
	@go test -v ./...