FROM golang:latest AS build

# Setting up working directory
WORKDIR $GOPATH/src/github.com/taufanmahaputra/ijah
COPY . .

# Install dependencies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# Build inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /ijah ./cmd/ijah/

# Single layer image
FROM scratch
COPY --from=build /ijah .
COPY ./config ./config
CMD ["/ijah"]