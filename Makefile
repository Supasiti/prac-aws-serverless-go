HANDLERS = "get_user" 



.PHONY: deploy
deploy:
	cdk deploy

.PHONY: build
build: clean
	for handler in $(HANDLERS) ; do \
		go build -o "./app/handler/$${handler}" "./cmd/handler/$${handler}.go" ; \
	done

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: run-user
run-user: build dep
	./app/user_dao

.PHONY: dep
dep:
	docker-compose up -d

.PHONY: clean
clean:
	rm -rf app/*

