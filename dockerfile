FROM golang:1.18.3-alpine3.16 as BUILD
COPY . /build
WORKDIR /build
RUN apk add build-base
RUN go build
FROM alpine:3.16 as DEPLOY
COPY --from=BUILD /build/go-product-api /
VOLUME [ "/data" ]
ENV DATA_DIR=/data
ENV GIN_MODE=release
CMD [ "/go-product-api" ]
EXPOSE 8080