# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.13.4 as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/josiemundi/knative-eventing-web-eventsource-server
COPY . .

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.4/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN dep init
RUN dep ensure -vendor-only


RUN CGO_ENABLED=0 GOOS=linux go build -v -o eventdisplay

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates

# Copy the binary and public html to the production image from the builder stage.
COPY --from=builder /go/src/github.com/josiemundi/knative-eventing-web-eventsource-server/eventdisplay /eventdisplay
COPY --from=builder /go/src/github.com/josiemundi/knative-eventing-web-eventsource-server/public /public

# Run the web service on container startup.
CMD ["/eventdisplay"]
