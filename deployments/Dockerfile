FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/l0

RUN CGO_ENABLED=0 GOOS=linux go build -o /app-build

EXPOSE 8080

CMD ["/app-build"]


