FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o workshop main.go
EXPOSE 8080

FROM scratch
COPY --from=builder /go/src/app/workshop /
CMD ["./workshop"]