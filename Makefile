service:
	go run ./...

mocks-clean:
	rm -rf mocks/*

mocks: mocks-clean mocks-generate

mocks-generate:
	mockery --all --output=mocks --case=underscore --keeptree

test:
	go test -v ./...