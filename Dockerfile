FROM alpine:3.7

RUN apk add --no-cache \
    ca-certificates \
    git \
    go \
    musl-dev \
    tzdata

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src/github.com/jlyon1/IEDFoodStorage" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/github.com/jlyon1/IEDFoodStorage

COPY . .

RUN go get github.com/gorilla/mux
RUN go get -v github.com/spf13/viper
RUN go get github.com/go-sql-driver/mysql
RUN go build -o iedfood

EXPOSE 8081

CMD ["./iedfood"]
