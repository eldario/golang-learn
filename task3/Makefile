 #!/usr/bin/make

help: ## Show this help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[32m%-11s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ### Run main.go
	go run main.go

build: ### Build
	go build .

test: ### Run tests
	go test ./pkg

bench: ### Bench tests
	go test -bench=. ./pkg/mapper/

lint: ## Lint
	 golangci-lint run
