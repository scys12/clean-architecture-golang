# Start from golang base image
FROM golang:1.14-alpine as builder


# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod .
COPY go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
# -a means force rebuilding of packages that already up-to-date
# -installsuffix means a suffix to use in then name of package installation directory in order to keep output separate from default builds.
# -o means forces build to write the resulting executable or object to the named output file or directory, instead of the default behavior
# in this case, building main file to the working dir inside container
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the config.env file
COPY --from=builder /app/main .
COPY --from=builder /app/config.env .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]