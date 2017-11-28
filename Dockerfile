FROM golang:1.9.2 as builder

WORKDIR /go/src/bitbucket.org/skibish/trashdiena/

RUN go get -u github.com/golang/dep/cmd/dep

COPY . .

RUN dep ensure \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo bitbucket.org/skibish/trashdiena/cmd/authorizer \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo bitbucket.org/skibish/trashdiena/cmd/loader \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo bitbucket.org/skibish/trashdiena/cmd/scheduler

FROM scratch
COPY --from=builder /go/src/bitbucket.org/skibish/trashdiena/authorizer .
COPY --from=builder /go/src/bitbucket.org/skibish/trashdiena/loader .
COPY --from=builder /go/src/bitbucket.org/skibish/trashdiena/scheduler .

EXPOSE 80
ENTRYPOINT [ "/authorizer" ]
