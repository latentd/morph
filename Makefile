.PHONY: cover clean

build:
	go build -o morph .

cover:
	go test -v -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

clean:
	rm morph cover*
