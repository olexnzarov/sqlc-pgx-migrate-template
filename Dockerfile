FROM golang:1.22.3

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o example-app cmd/example-app/main.go

CMD ["./example-app"]
