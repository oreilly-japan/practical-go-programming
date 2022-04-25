FROM golang:1.17-bullseye as builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -ldflags="-w -s" -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=builder /go/bin/app /
CMD ["/app"]
