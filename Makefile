default: linux

output=web-server

linux: clean
	GOOS=linux GOARCH=amd64 \
	go build -o ${output} main.go

clean:
	rm -f ${output}
