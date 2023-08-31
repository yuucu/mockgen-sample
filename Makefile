
install:
	go install go.uber.org/mock/mockgen@latest

gen-mock:
	mockgen -source=random/random.go -destination=random/random_mock.go -package=random

run:
	go run main.go

test:
	go test -v ./...

