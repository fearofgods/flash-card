# Define the output binary name
BINARY_NAME=app.exe

run:
	go run .\cmd\main.go -e development

build:
	go build -o .\build\app.exe .\cmd\main.go 

swag:
	swag init --dir .\cmd -o .\docs

.PHONY: run build swag