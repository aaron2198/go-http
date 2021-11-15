FROM golang:latest

WORKDIR /src/go/app

RUN go get github.com/cespare/reflex
COPY . .
CMD [ "reflex", "-v", "-s", "-r", "\\.go$", "--", "go", "run", "."]