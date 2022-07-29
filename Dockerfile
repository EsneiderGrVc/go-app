FROM golang:1.18.4

WORKDIR /usr/src/app
RUN go mod init github.com/EsneiderGrVc/go_server
COPY ./api /usr/src/app
RUN go get
RUN go get github.com/githubnemo/CompileDaemon \
&& go install github.com/githubnemo/CompileDaemon
EXPOSE 80
