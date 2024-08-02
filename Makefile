build:
	@go build -o ./bin/bookbrowse
run:build
	@./bin/bookbrowse

test_stores:
	@go test -v ./stores/

test: test_stores

templ:
	@templ generate --watch --proxy=http://localhost:7000

css:
	@tailwindcss -i ./views/css/app.css -o ./public/styles.css --watch