FROM golang:alpine AS build

ENV GOPATH /go
WORKDIR /go/src
COPY . /go/src
RUN cd /go/src && go build .

FROM alpine
WORKDIR /app
COPY --from=build go/src/pismo /app
COPY .env /app

EXPOSE 8080

ENTRYPOINT [ "./pismo" ]