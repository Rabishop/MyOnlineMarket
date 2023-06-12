FROM golang:1.19.2 as builder

ENV MYPATH /usr/local
WORKDIR $MYPATH

ADD ./pages ./pages
ADD ./src ./src

WORKDIR /usr/local/src
RUN go build -o my-online-market-app main.go

EXPOSE 8080

# CMD ./my-online-market-app
