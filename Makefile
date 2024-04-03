compile:
	echo "Compiling for linux and windows"
	GOOS=linux GOARCH=amd64 go build -o bin/shoot-run-linux-x64 main.go
	GOOS=linux GOARCH=386 go build -o bin/shoot-run-linux-x86 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/shoot-run-windows-x64.exe main.go
	GOOS=windows GOARCH=386 go build -o bin/shoot-run-windows-x86.exe main.go

run:
	go run main.go

tests:
	go test ./game/tests -v