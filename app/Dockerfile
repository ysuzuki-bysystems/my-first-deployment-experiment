FROM golang:1.23.1-bookworm AS builder

RUN mkdir /build
WORKDIR /build
COPY ./go.mod ./main.go .
RUN CGO_ENABLED=0 go build -o app .

FROM scratch
COPY --from=builder /build/app /app
CMD ["/app"]
