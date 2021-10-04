FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go get ./... 

RUN go build -o ./bin/web-app ./cmd/apiserver/main.go
ENTRYPOINT /app/bin/web-app
