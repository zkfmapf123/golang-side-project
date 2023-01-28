FROM golang:1.18.0

WORKDIR /usr/src/app

RUN export GO111MODULE=on

COPY go.mod ./

RUN go mod download

COPY . .

# Build the application
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
