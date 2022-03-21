FROM golang:1.16 as builder
WORKDIR /app
COPY . .
EXPOSE 9090
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./serviceapp ./cmd/app/main.go

# Up binary
FROM scratch
COPY --from=builder /app/serviceapp /usr/bin/serviceapp
ENTRYPOINT ["/usr/bin/serviceapp" ]