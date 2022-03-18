# Start from the latest golang base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY ./server .
COPY ./server/migration ./migration


RUN go mod tidy

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migration ./migration

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/migration/migration .
COPY --from=builder /app/migration/scripts/*.sql ./scripts/

# Expose port 8080 to the outside world
EXPOSE 8080