FROM golang:1.9.2 as builder

WORKDIR /go/src/bitbucket.org/skibish/trashdiena/

RUN go get -u github.com/golang/dep/cmd/dep

COPY . .

RUN dep ensure && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o trashdiena .


FROM scratch
COPY --from=builder /go/src/bitbucket.org/skibish/trashdiena/trashdiena .
ENTRYPOINT [ "/trashdiena" ]
