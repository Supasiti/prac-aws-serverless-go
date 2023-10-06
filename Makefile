.PHONY: build
build:
	go build -o ./app/user_dao ./cmd/user_dao.go

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
	docker-compose down -v

