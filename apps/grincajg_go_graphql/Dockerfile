# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM --platform=linux/amd64 golang:1.21.3-bullseye as builder

# Copy local code to the container image.
WORKDIR /app
COPY . .

# Build the command inside the container.
# (You may need to modify this line if you have specific build requirements.)
RUN go build -v -o grincajg_go

# Use the official Debian slim image for a smaller image.
FROM --platform=linux/amd64 debian:buster-slim

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/grincajg_go /grincajg_go

COPY .env /

EXPOSE 8090

# Run the web service on container startup.
CMD ["/grincajg_go"]
