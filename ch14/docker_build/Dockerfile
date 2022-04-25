FROM golang:1.17-bullseye as builder
WORKDIR /opt/app
COPY go.mod go.sum ./
RUN go mod download
COPY ./*.go ./
COPY ./cmd/sample/*.go ./cmd/sample/
RUN cd cmd/sample && go build -trimpath -ldflags="-w -s" -o "sample-app"

FROM gcr.io/distroless/base-debian11
COPY --from=builder /opt/app/cmd/sample/sample-app /sample-app
CMD ["/sample-app"]
