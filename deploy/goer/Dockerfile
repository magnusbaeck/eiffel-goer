FROM golang:1.16.3-alpine AS build
WORKDIR /tmp/goer
COPY . .
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go build -o ./bin/goer ./cmd/goer

FROM alpine:3.13.5
COPY --from=build /tmp/goer/bin/goer /app/goer
CMD ["/app/goer"]