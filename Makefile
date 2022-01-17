build:
	go build -o bin/log app/log.go

run:
	go run app/log.go


compile:
	echo "Compiling for linux mac and win"
	GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/log app/log.go
	GOOS=linux GOARCH=386 go build -o bin/linux/386/log app/log.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/log app/log.go
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/amd64/log app/log.go
	GOOS=windows GOARCH=amd64 go build -o bin/win/amd64/log app/log.go
	GOOS=windows GOARCH=386 go build -o bin/win/386/log app/log.go