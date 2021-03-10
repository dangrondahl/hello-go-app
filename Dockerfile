# Step 1 - build the executable
FROM golang:alpine

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# Copy the local package to the container's workspace.
COPY . $GOPATH/src/dangrondahl/argo-hello-go-app/
WORKDIR $GOPATH/src/dangrondahl/argo-hello-go-app/

# Get dependencies
#RUN go get -d -v

# Build the binary
RUN go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/helloworld

