 #!/usr/bin/make

help: ## Show this help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[32m%-11s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ### Run main.go
	go run main.go

build: ### Build
	go build .

lint: ## Lint
	 golangci-lint run

cpu: ## CPU Pprof
	go tool pprof cpu.prof

mem: ## MEM Pprof
	go tool pprof mem.prof
