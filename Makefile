windows:
	GOOS=windows go build -o pkgpal.exe

mac:
	GOOS=darwin go build -o pkgpal

linux:
	GOOS=linux go build -o pkgpal

bin:
	sudo mv pkgpal /usr/local/bin/

setup:
	go get -t ./...

test:
	go test ./...