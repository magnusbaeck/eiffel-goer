FROM golang:1.21.3-alpine AS build

RUN apk update && apk add --no-cache git make

WORKDIR /tmp/goer
COPY . .
RUN make build

FROM alpine:3.19.0
ENTRYPOINT ["/app/goer"]
COPY --from=build /tmp/goer/bin/goer /app/goer

ARG revision
LABEL com.datadoghq.tags.version="${revision}" \
      org.opencontainers.image.licenses="Apache-2.0" \
      org.opencontainers.image.revision="${revision}" \
      org.opencontainers.image.source="https://github.com/eiffel-community/eiffel-goer"
