# backend/Dockerfile
FROM golang:1.23-alpine

WORKDIR /app


COPY go.mod go.sum ./
# Install air for live reloading

ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
RUN go mod download


COPY . .

RUN go build -o main .
RUN go install github.com/cosmtrek/air@v1.29.0

EXPOSE 8080

CMD ["./main"]