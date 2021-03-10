# Step 1 - build the executable
FROM golang:alpine as builder

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# Copy the local package to the container's workspace.
WORKDIR /tmp/hello-app/

COPY . .

# Get dependencies
RUN go get -d -v

# Build the binary
RUN go build -o ./out/hello .

# Step 2 - Build a smaller image
FROM scratch

COPY --from=builder /tmp/hello-app/out/hello /app/hello

ENTRYPOINT [ "/app/hello" ]

EXPOSE 8080
