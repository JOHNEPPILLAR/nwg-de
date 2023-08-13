# syntax=docker/dockerfile:1

# Base image
FROM golang:alpine as base-build
WORKDIR /app

# Update and install git
RUN apk update && apk add git

# Copy go install files
COPY go.mod ./ 
COPY go.sum ./

# Download packages
RUN go mod download && go mod verify

# Build app
COPY . .
RUN go build -o /nwg-de cmd/main.go

# Start fresh from a smaller image
FROM alpine:latest

# Add certs and tzdata
RUN apk update && apk add ca-certificates && apk add tzdata

# Set timezone and copy built app
ENV TZ=Europe/London
COPY --from=base-build /nwg-de ./nwg-de

# Run binary
ENV PORT 8080
EXPOSE 8080
CMD ["./nwg-de"]