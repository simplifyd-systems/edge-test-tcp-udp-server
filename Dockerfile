FROM golang:1.17 as builder

# Add metadata identifying these images as our build containers 
# (this will be useful to delete later to prevent leaking sensitive information!)
LABEL stage=intermediate

# It is important that these ARG's are defined after the FROM statement
ARG GITHUB_PERSONAL_ACCESS_TOKEN="nothing"

RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates

RUN update-ca-certificates

RUN printf "machine github.com login ${GITHUB_PERSONAL_ACCESS_TOKEN}\n"\
    >> /root/.netrc
RUN chmod 600 /root/.netrc

# Move to working directory /build
WORKDIR /build

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy and download dependency using go mod
COPY go.mod .
# COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

FROM scratch as final

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /build/main /go/bin/

ENTRYPOINT ["/go/bin/main"]