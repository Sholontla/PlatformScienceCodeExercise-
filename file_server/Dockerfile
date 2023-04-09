FROM golang:1.19

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ${ADMIN_DIR} .

CMD ["go", "run", "/app/cmd/application/main.go"]