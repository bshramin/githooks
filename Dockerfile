
FROM golang:1.18.3 as build

RUN apt update
RUN apt install git

ENV BIN=/go/src/github.com/bshramin/githooks/bin
ENV GOBIN=/go/src/github.com/bshramin/githooks/bin

COPY . /go/src/github.com/bshramin/githooks
WORKDIR /go/src/github.com/bshramin/githooks

RUN go install -mod=vendor -v -ldflags "-s" ./cmd/...

# -----------------------------------------------------------------------------
FROM golang:1.18.3

COPY --from=build /go/src/github.com/bshramin/githooks/bin ./bin

EXPOSE 8008

CMD ["bin/gitlab"]
