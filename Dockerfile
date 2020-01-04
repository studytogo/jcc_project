FROM golang:latest
  
MAINTAINER jcc

ENV TZ Asia/Shanghai

RUN echo 'Asia/Shanghai' >/etc/timezone

RUN go get github.com/beego/bee

#ENV GOPROXY https://goproxy.io/

#ENV GO111MODULE on

WORKDIR $GOPATH/src/new_erp_agent_by_go

ADD . .

#RUN go mod vendor

EXPOSE 8080

ENV GO111MODULE off

#CMD bee run -gendoc=true

RUN chmod +x ./run.sh

CMD ./run.sh