FROM golang:1.17.7-bullseye AS go-builder
WORKDIR /source
COPY Makefile Makefile
COPY go.mod go.mod
COPY go.sum go.sum
RUN make dep-go
COPY . .
RUN make build-go

FROM python:3.9-bullseye
WORKDIR /py-dep
COPY Makefile Makefile
COPY requirements.txt requirements.txt
RUN make dep-py
WORKDIR /app
RUN rm -rf /py-dep
COPY --from=go-builder /source/bin/hn-api ./hn-api

CMD ["/app/hn-api"]