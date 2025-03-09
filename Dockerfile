FROM golang:1.23.7-alpine3.21

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./cmd/th3-sh0p-api-server/main.go
EXPOSE 80

CMD ["app", "--host", "0.0.0.0", "--port", "80"]