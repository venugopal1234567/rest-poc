FROM golang:1.16.5-alpine3.13 AS build

WORKDIR /go/src/rest-poc
COPY . .

RUN go build -mod=vendor main.go

FROM alpine:3.13.0

COPY --from=build /go/src/rest-poc/main ./bin/main

EXPOSE 8000

CMD ["./bin/main"]