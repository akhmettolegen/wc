run: ### run
	go mod tidy && \
	go run .

test: ### run test
	go test -v ./...
