FROM golang:1.18 as dev

WORKDIR WORKDIR $GOPATH/src/ports

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8000
CMD ["go", "run", "main.go"]
