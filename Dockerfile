# Step 1 - build the executable
FROM golang:1.16-alpine as builder

ARG PROJECT="github.com/dangrondahl/hello-go-app"
ARG VERSION
ARG COMMIT
ARG BUILD_TIME

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# Copy the local package to the container's workspace.
WORKDIR /tmp/hello-app/

COPY . .

# Get dependencies
RUN go get -d -v

# Build the binary
RUN go build \
  -ldflags "-s -w -X ${PROJECT}/version.Version=${VERSION} \
  -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
  -o ./out/hello .


# Run unit tests
#RUN go test -v ./...

# Step 2 - Build a smaller image
FROM scratch

COPY --from=builder /tmp/hello-app/out/hello /app/hello

ENTRYPOINT [ "/app/hello" ]

EXPOSE 8080
