FROM golang:1.19.2 as builder

ENV MYPATH /usr/local
WORKDIR $MYPATH

ADD ./pages ./pages
ADD ./src ./src

EXPOSE 8080

CMD ./src/my-online-market
