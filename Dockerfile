# Define argument for linker flags
ARG LDFLAGS=-s -w

# Use a temporary build image based on Golang 1.20-alpine
FROM golang:1.24-alpine as builder

# Set environment variables: linker flags and disable CGO
ENV LDFLAGS=$LDFLAGS CGO_ENABLED=0

# Add the current directory contents to the work directory in the container
ADD . /work

# Set the current work directory inside the container
WORKDIR /work

# Install git and build the WarpNet binary with the provided linker flags
# --no-cache flag ensures the package cache isn't stored in the layer, reducing image size
RUN apk add --no-cache git && \
    go build -ldflags="$LDFLAGS" -o WarpNet

# TODO: move to distroless

# Use a new, clean alpine image for the final stage
FROM alpine

# Copy the WarpNet binary from the builder stage to the final image
COPY --from=builder /work/WarpNet /usr/bin/WarpNet

# Define the command that will be run when the container is started
ENTRYPOINT ["/usr/bin/WarpNet"]
