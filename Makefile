run:
	go run main.go

run-docker:
	docker-compose up
run-tests:
	go test -v
	go test -coverprofile=coverage.out
	go tool cover -html=c.out -o coverage.html

run-populate: run-docker
	go run cmd/migration/main.go