FROM golang:1.13.3-stretch

LABEL maintainer="Devon Tingley <dtingley@twilit.io>"

WORKDIR /usr/source/app
COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN go build -o api

EXPOSE 8080
CMD ["./api"]