FROM golang:alpine AS builder

RUN apk update && \
    apk add git && \
    apk add build-base upx

WORKDIR /src/app
COPY . .
RUN go build  -o /go/bin/logstealer cmd/logstealer/main.go
RUN upx /go/bin/logstealer

FROM alpine:latest
#RUN apk update && apk add --no-cache  vips-dev
COPY --from=builder /go/bin/logstealer /go/bin/logstealer

ENTRYPOINT ["/go/bin/logstealer"]

