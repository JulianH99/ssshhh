
run:
	@go run cmd/main.go

compile:
	@rm build/ssshhh.exe
	@go build -o build/ssshhh.exe github.com/JulianH99/ssshhh/cmd
