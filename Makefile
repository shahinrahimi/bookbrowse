build:
	@go build -o ./bin/bookbrowse
run:build
	@./bin/bookbrowse

test_stores:
	@go test -v ./stores/

test: test_stores