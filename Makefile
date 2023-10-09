HANDLERS = "get_user" 

.PHONY: deploy
deploy: tidy
	cdk deploy

.PHONY: build
build: tidy clean
	for handler in $(HANDLERS) ; do \
		go build -o "./app/handler/$${handler}" "./cmd/handler/$${handler}.go" ; \
	done

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -rf app/*

