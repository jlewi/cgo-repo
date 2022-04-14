build:
	mkdir -p .build
	CGO_ENABLED=0 go build -o .build/nocgo .
	CGO_ENABLED=1 go build -o .build/cgo .
