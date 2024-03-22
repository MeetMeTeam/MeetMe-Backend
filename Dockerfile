FROM golang:1.18-alpine as builder

WORKDIR /builder

# Install UPX for later use
RUN apk add --no-cache upx

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0

# Cache Go modules separately to speed up subsequent builds
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application with UPX compression
RUN go build -ldflags "-s -w" -o /builder/main /builder/main.go \
    && upx -9 /builder/main


# Runner stage
FROM gcr.io/distroless/static:latest

WORKDIR /app

# Copy the binary and HTML templates from the builder stage
COPY --from=builder /builder/main .
COPY reset-password.html .
COPY verify-mail.html .

# Expose the port
EXPOSE 8080

# Define the command to run the application
CMD ["./main"]
