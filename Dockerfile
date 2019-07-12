FROM golang:alpine
ADD . ${HOME}/merge-tool
WORKDIR ${HOME}/merge-tool
RUN go build -o main . 

VOLUME /files/reports
VOLUME /files/result

CMD ["./main"]