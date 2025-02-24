FROM golang:1.22

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
COPY . .
RUN go mod download

CMD ["air", "-c", ".air.toml"]