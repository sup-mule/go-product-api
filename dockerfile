FROM golang:1.20.13-alpine3.18 as BUILD
COPY . /build
WORKDIR /build
RUN apk add build-base
RUN go build
FROM alpine:3.18 as DEPLOY
COPY --from=BUILD /build/go-product-api /
VOLUME [ "/data" ]
ENV DATA_DIR=/data
ENV GIN_MODE=release
CMD [ "/go-product-api" ]
EXPOSE 8080