FROM golang:1.17.7-alpine as build

ARG BUILD_TARGET

WORKDIR /${BUILD_TARGET}

RUN set -eux \
        && apk update \
        && apk add alpine-sdk \
        && go install github.com/cosmtrek/air@v1.27.4

ADD client /client
ADD ${BUILD_TARGET}/go.mod .
ADD ${BUILD_TARGET}/go.sum .
RUN go mod download

ADD ${BUILD_TARGET} .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags="all=-N -l" -o /app/main

COPY .env /app/.env

FROM alpine:latest as prod

COPY --from=build /etc/localtime /etc/localtime
COPY --from=build /app/main /app/main
COPY --from=build /app/.env /app/.env

EXPOSE 80
ENTRYPOINT ["/app/main"]