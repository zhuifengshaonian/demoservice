FROM golang:1.18 AS builder

WORKDIR /go/src/app
COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://repo.huaweicloud.com/repository/goproxy/

RUN mkdir bin
RUN go build -o ./bin ./...

FROM ubuntu:rolling
COPY --from=builder /go/src/app/bin /app
WORKDIR /app
VOLUME /conf
VOLUME /logs
EXPOSE 9000
EXPOSE 9001
ENTRYPOINT ["./demoservice"]