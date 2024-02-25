FROM golang:alpine AS builder

RUN apk update && \
    apk add git && \
    apk add build-base upx

WORKDIR /src/app
COPY . .
RUN go build  -o /go/bin/app cmd/logstealer/main.go
RUN upx /go/bin/app

FROM alpine
#RUN apk update && apk add --no-cache  vips-dev
COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]

