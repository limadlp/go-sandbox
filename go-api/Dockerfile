# Build Stage
FROM golang:1.23 AS builder

WORKDIR /app

# Copy only the necessary files for building
COPY go.mod go.sum ./
RUN go mod download

# Copy the remaining code
COPY . .

# Compile the binary statically to avoid dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Runtime Stage (final image)
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary
COPY --from=builder /app/main .

# Give execution permission to the binary (if needed)
RUN chmod +x ./main

# Expose the port used by the app
EXPOSE 8000

# Startup command
CMD ["./main"]


## run: docker compose up --build -d