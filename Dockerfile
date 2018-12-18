FROM golang:alpine
ENV PROJ_DIR=$GOPATH/src/bitbucket.org/hameesys/merge-xml-report
RUN mkdir -p ${PROJ_DIR}
ADD . ${PROJ_DIR}
WORKDIR ${PROJ_DIR}
RUN go build -o main . 

VOLUME /files/reports
VOLUME /files/result

CMD ["./main"]