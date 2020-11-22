 #!/usr/bin/make

help: ## Show this help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[32m%-11s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ### Run main.go
	go run main.go

build: ### Build
	go build .

test: ### Run tests
	go test ./pkg/sorter/sorter_test.go && go test ./pkg/sorter/list_test.go