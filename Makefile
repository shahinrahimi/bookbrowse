build:
	@go build -o ./bin/bookbrowse
run:build
	@./bin/bookbrowse

test_store:
	@go test -v ./store/

test: test_store