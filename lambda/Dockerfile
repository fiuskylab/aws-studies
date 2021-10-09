FROM golang:1.17 as build

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN GOOS=linux GOARCH=amd64 go build -o main main.go

CMD ["./main"]
