FROM golang:1.18-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod verify

COPY ["pkg", "/app/pkg/"]
COPY ["cmd", "/app/cmd/"]

RUN go build -o judger judger/cmd/judger

FROM alpine:latest as run
WORKDIR /app
COPY --from=docker:dind /usr/local/bin/docker /usr/local/bin/
COPY --from=build /app/judger /app/judger
CMD ["./judger"]
