HANDLERS = "get_user" 

.PHONY: deploy
deploy: tidy
	cdk deploy

.PHONY: build
build: tidy clean
	for handler in $(HANDLERS) ; do \
		go build -o "./app/handler/$${handler}" "./cmd/handler/$${handler}/main.go" ; \
	done

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -rf app/*

.PHONY: test
test: 
	go test -short -race -gcflags=all=-l ./...

.PHONY: test-one
test-one:
	go test -short -race -gcflags=all=-l $(path) -v -run $(fn) 
	
.PHONY: generate
generate:
	go generate ./...

.PHONY: lint
lint: deps
	bin/golangci-lint run

.PHONY: fmt
fmt: deps 
	bin/golangci-lint run --fix

deps: bin/golangci-lint
deps: # install dependencies

GOLANGCI_VERSION = 1.53.3

bin/golangci-lint:
	@mkdir -p bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- v${GOLANGCI_VERSION}
