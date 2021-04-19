build:
	go mod tidy
	go build -o bin/binary .
	bin/binary
	ls -l bin/binary