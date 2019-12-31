FROM golang:latest

RUN apt-get update && apt-get upgrade
RUN apt-get install curl -y

WORKDIR /go/src/github.com/barokurniawan/YoutubeDownloader

COPY ./ /go/src/github.com/barokurniawan/YoutubeDownloader

RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl

RUN chmod a+rx /usr/local/bin/youtube-dl

RUN go get -u github.com/gorilla/mux

ENTRYPOINT go run main.go