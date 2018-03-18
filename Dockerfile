FROM        golang:onbuild

WORKDIR /go/src/github.com/alexanderbrunner/docker-go-examples/mygoservice
COPY mygoservice .

RUN go get -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD["mygoservice"]
