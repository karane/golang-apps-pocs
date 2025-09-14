FROM golang:1.23-alpine

WORKDIR /app

# Install SQLite (needed for migration in entrypoint.sh)
RUN apk add --no-cache sqlite
# Install build dependencies for cgo + SQLite
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Copy project files
COPY . .


# Enable CGO
ENV CGO_ENABLED=1

# Initialize Go module (if not present) and download dependencies
RUN go mod init url-shortener || true
RUN go mod tidy

# Build the Go binary
RUN go build -o url-shortener main.go

# Make entrypoint executable
RUN chmod +x entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]
# CMD ["tail","-f", "/dev/null"]