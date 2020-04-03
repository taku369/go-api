GO111MODULE := on

all:
	go get -u github.com/gorilla/mux
	go build .

run:
	go run .

clean:
	go clean .
