# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/sarathsp06/gorest

# Get dependencies
RUN go get -v  github.com/sarathsp06/gorest

WORKDIR /go/src/github.com/sarathsp06/gorest
RUN make compile
RUN make setup

# Run the service
ENTRYPOINT ["make","run"]

# Document that the service listens on port 8080.
EXPOSE 8080