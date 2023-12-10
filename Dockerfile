FROM golang:1.21.1

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"

CMD ["tail", "-f", "/dev/null"]