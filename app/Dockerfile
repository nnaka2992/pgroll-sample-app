FROM golang:1.21.5-alpine3.19 AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11
COPY --from=builder /go/bin/app /app
CMD ["/app"]
EXPOSE 8080
