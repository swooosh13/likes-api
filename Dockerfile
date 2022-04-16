FROM golang:1.18 as builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./serviceapp ./cmd/app/main.go
ENTRYPOINT ["/app/serviceapp" ]

# Up binary
# FROM scratch
# COPY --from=builder /app/serviceapp /usr/bin/serviceapp
# ENTRYPOINT ["/usr/bin/serviceapp" ]