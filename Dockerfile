# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /cp-changelogs

FROM alpine:3.10

COPY --from=build /cp-changelogs /cp-changelogs

EXPOSE 80

CMD [ "/cp-changelogs" ]