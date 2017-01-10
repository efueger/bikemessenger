package = github.com/delivercodes/bikemessenger

.PHONY: release

release:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/bikemessenger-linux-amd64 $(package)
	GOOS=linux GOARCH=386 go build -o release/bikemessenger-linux-386 $(package)
	GOOS=linux GOARCH=arm go build -o release/bikemessenger-linux-arm $(package)
