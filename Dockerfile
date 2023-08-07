# Build Golang binary
FROM golang:1.20.2 AS build-golang

WORKDIR D:\Projects\Go\learn-beego

COPY . .
RUN go get -v && go build -v -o /usr/local/bin/D:\Projects\Go\learn-beego

EXPOSE 8080
CMD ["D:\Projects\Go\learn-beego"]
