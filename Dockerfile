# Stage build
FROM golang:1.15.6-alpine AS builder
ENV GO111MODULE=on

RUN apk add --update gcc openssh git bash libc-dev ca-certificates make g++
ENV BUILDDIR /go/src/game-currency

ENV ENV production
COPY . /go/src/game-currency
WORKDIR /go/src/game-currency
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /go/src/game-currency/game-currency /go/src/game-currency/main.go

# Stage Runtime Applications
FROM alpine:latest

# Setting timezone
ENV TZ=Asia/Jakarta
RUN apk add -U tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# add ca-certificates
RUN apk add --no-cache ca-certificates

ENV BUILDDIR /go/src/game-currency

# Setting folder workdir
WORKDIR /opt/
RUN mkdir files
RUN mkdir files/etc
# Copy Data App
COPY --from=builder $BUILDDIR/game-currency .
COPY --from=builder $BUILDDIR/files/etc/main.development.ini /opt/files/etc/main.development.ini

EXPOSE 8081

ENTRYPOINT ["./game-currency"]