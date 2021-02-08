FROM golang:1.14-alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/server cmd/server/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /build/bin/server .
EXPOSE 3000
CMD [ "/app/server" ]
