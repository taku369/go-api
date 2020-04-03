FROM golang:latest

WORKDIR /app

COPY app.go .
COPY Makefile .

RUN make

CMD ["./app"]
